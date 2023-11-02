package main

import (
	"log"

	"github.com/fishmanDK/internal/handlers"
	postgresql "github.com/fishmanDK/internal/repository/postgreSQL"
	"github.com/fishmanDK/internal/service"
)


func main() {
	// serv_config := configs.NewServer()
	// db_config := configs.NewPostgreSQL()
	
	postgre_cfg := postgresql.InitPostgreConfig()
	repo, err := postgresql.NewPostgreDB(postgre_cfg)
	if err != nil{
		log.Fatal(err)
	}
	service := service.NewService(repo)
	handlers := handlers.NewHandlers(service)
	
	serv := handlers.InitRouts()

	serv.Run(":8080")
}