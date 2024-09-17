package container

import "github.com/google/wire"

type Container struct {
	service int
}

func InitContainer(db interface{}) {
	wire.Build()
}
