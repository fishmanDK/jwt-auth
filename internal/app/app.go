package app

import (
	grpcapp "github.com/fishmanDK/internal/app/grpc"
	"github.com/fishmanDK/internal/repository"
	"github.com/fishmanDK/internal/service"
	"log"

	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(logger *slog.Logger, storage, port string) *App {
	db, err := repository.NewStorage(storage)
	if err != nil {
		log.Fatalf("No connection to %s", storage)
	}

	srvc := service.NewService(db)

	app := grpcapp.NewApp(logger, srvc, port)

	return &App{
		GRPCSrv: app,
	}
}
