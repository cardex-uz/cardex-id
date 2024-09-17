package app

import (
	implapp "cardex-id/internal/app/grpc"
	"cardex-id/internal/container"
	"log/slog"
)

type App struct {
	Server *implapp.App
}

func New(log *slog.Logger, port int) *App {
	storage := ""
	container := container.InitContainer(storage)
}
