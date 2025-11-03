package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"

	authgrpc "github.com/Numbone/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Reqister(gRPCServer)
	return &App{
		log:  log,
		gRPC: gRPCServer,
		port: port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcApp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("port", strconv.Itoa(a.port)),
	)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("starting grpc server", slog.String("addr", lis.Addr().String()))
	if err := a.gRPC.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcApp.Stop"
	a.log.With(slog.String("op", op)).
		Info("stopping grpc server", slog.Int("port", a.port))

	a.gRPC.GracefulStop()
}
