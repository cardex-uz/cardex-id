package container

import (
	dUsecase "cardex-id/internal/domain/usecase"
	"gorm.io/gorm"
)

type DB = gorm.DB

type Container struct {
	service dUsecase.Auth
}

func Init(db *DB) *Container {
	return &Container{
		//service: usecase.NewAuth(),
	}
}
