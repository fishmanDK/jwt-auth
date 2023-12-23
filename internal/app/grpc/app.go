package grpc

import (
	"fmt"
	"github.com/fishmanDK/internal/handlers/h_grpc"
	"github.com/fishmanDK/internal/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	Logger     *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func NewApp(logger *slog.Logger, service *service.Service, port string) *App {
	s := grpc.NewServer()
	h_grpc.Register(s, service)

	return &App{
		Logger:     logger,
		gRPCServer: s,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.app.Run"

	log := a.Logger.With(
		slog.String("op", op),
		slog.String("port", a.port),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", lis.Addr().String()))

	if err = a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpc.app.Run"

	a.Logger.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
