//go:build wireinject
// +build wireinject

package main

import (
	"wire/demo"

	"github.com/google/wire"
)

func initZ() (demo.Z, error) {
	// 不使用wire、手动New
	// x := demo.NewX()
	// y := demo.NewY(x)
	// z, err := demo.NewZ(y)
	// fmt.Println(z, err)

	// 使用wire
	// wire.Build(demo.NewX, demo.NewY, demo.NewZ)
	/*
		a call to wire.Build indicates that this function is an injector,
		but injectors must consist of only the wire.Build call and an optional return
	*/
	wire.Build(demo.ProviderSet) // but injectors must consist of only the wire.Build call and an optional return
	return demo.Z{}, nil
}
