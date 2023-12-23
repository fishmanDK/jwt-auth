package main

import (
	"github.com/fishmanDK/configs"
	app "github.com/fishmanDK/internal/app"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	postgres = "postgres"
	port     = "5001"
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := configs.NewConfig()
	logger := setupLogger(cfg.Env)

	appgrpc := app.NewApp(logger, postgres, port)

	go func() {
		appgrpc.GRPCSrv.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	appgrpc.GRPCSrv.Stop()

	//logger.Info("application stopped", slog.String("signal", signalStop.String()))
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
