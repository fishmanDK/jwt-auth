package main

import (
	"fmt"
	"github.com/fishmanDK/configs"
	http2 "github.com/fishmanDK/internal/handlers/h_http"
	"github.com/fishmanDK/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/fishmanDK/internal/service"
)

const (
	postgres = "postgres"
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := configs.NewConfig()
	logger := setupLogger(cfg.Env)

	fmt.Println(cfg)

	db, err := repository.NewStorage(postgres)
	if err != nil {
		log.Fatalf("No connection to %s", postgres)
	}
	service := service.NewService(db)
	handlers := http2.NewHandlers(service)
	routs := handlers.InitRouts(logger)

	server := initServer(cfg, routs)
	server.ListenAndServe()
}

func initServer(cfg *configs.Config, routs *gin.Engine) *http.Server {
	return &http.Server{
		Addr:              cfg.Address,
		Handler:           routs,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler := slog.NewTextHandler(os.Stdout, opts)
		logger = slog.New(handler)
	case envDev:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(handler)
	default:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(handler)
	}

	return logger
}
