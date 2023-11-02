package postgresql

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	jwtauth "github.com/fishmanDK"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
)

type Postgre interface {
	CreateUser(new_user jwtauth.CreateUser) error
	Authentication(user jwtauth.User) (AuthResult, error)
	AddRefresh(refresh_token string, expiresAt time.Time) error
}

type PostgreDB struct {
	postgre_db *sqlx.DB
}

func NewPostgreDB(cfg postgreConfig) (*PostgreDB, error) {
	link := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", link)
	
	if err != nil {
		return nil, err
	}
	fmt.Println(1)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreDB{
		postgre_db: db,
	}, nil
}

type postgreConfig struct {
	DBName   string `yaml:"name_db"`
	User     string `yaml:"user_db"`
	Password string `yaml:"password_db"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

func InitPostgreConfig() postgreConfig {
	config_path := "internal/repository/postgreSQL/postgre.yaml"
	if config_path == "" {
		log.Fatal("config_path is not set")
	}

	if _, err := os.Stat(config_path); err != nil {
		log.Fatalf("config file does not exist: %s", config_path)
	}

	var cfg postgreConfig

	if err := cleanenv.ReadConfig(config_path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err.Error())
	}
	return cfg
}
