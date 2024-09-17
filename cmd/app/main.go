package main

import (
	"cardex-id/internal/config"
	"cardex-id/pkg/logger/slogpretty"
	"cardex-id/pkg/postgresql"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type User struct {
	gorm.Model
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	db, err := postgresql.New(log, &cfg.DB)

	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	if err := db.AutoMigrate(); err != nil {
		fmt.Errorf("%s", err.Error())
	}

	//container := Container
	//application := app.New(log, cfg.Server.Port, cfg.StoragePath, cfg.TokenTTL)
	//
	//go func() {
	//	application.Server.MustRun()
	//}()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
