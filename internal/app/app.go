package app

import (
	"log/slog"
	"time"

	"github.com/Numbone/sso/internal/app/grpc"
)

type App struct {
	GRPCSrc *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	gprcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrc: gprcApp,
	}
}
