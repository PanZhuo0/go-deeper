package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := biz()
	fmt.Printf("service:%+v\n", err)
}

func dao() error {
	return nil
}

type user struct {
	size int
}

var ErrSizeTooSmall = errors.New("尺寸过小")

func biz() error {
	u := &user{size: 30}
	err := dao()
	if err != nil {
		return err
	}
	if u.size <= 30 {
		return errors.Wrap(ErrSizeTooSmall, "Error:size small!")
	}
	return nil
}
