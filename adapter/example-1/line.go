package main

import "fmt"

type Line struct{}

func (l *Line) draw(x, y, z, j int) {
	fmt.Println("Line - ", x, y, z, j)
}
