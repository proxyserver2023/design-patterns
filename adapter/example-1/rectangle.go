package main

import (
	"fmt"
)

type Rectangle struct{}

func (r *Rectangle) draw(x, y, z, j int) {
	fmt.Println("Rectangle - ", x, y, z, j)
}
