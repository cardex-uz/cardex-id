package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func New(log *slog.Logger, cfg *Config) (*gorm.DB, error) {
	const op = "data.Storage"

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Name,
		cfg.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Postgresql database started!")

	if err := db.AutoMigrate(); err != nil {
		fmt.Errorf("%s", err.Error())
	}

	return db, nil
}
