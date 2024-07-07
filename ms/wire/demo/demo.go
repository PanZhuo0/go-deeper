package demo

import (
	"errors"

	"github.com/google/wire"
)

type X struct {
	Value int
}

// 这是Wire中的一个Provider
func NewX() X {
	return X{Value: 1}
}

type Y struct {
	Value int
}

// 这也可以是Wire中的一个Provider
func NewY(x X) Y {
	return Y{Value: x.Value + 1}
}

type Z struct {
	Value int
}

// 这也是一个Wire中的Provider
func NewZ(y Y) (Z, error) {
	if y.Value == 0 {
		return Z{}, errors.New("invalid y")
	}
	return Z{Value: y.Value + 2}, nil
}

var ProviderSet = wire.NewSet(NewX, NewY, NewZ)
