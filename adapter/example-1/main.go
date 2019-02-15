package main

func main() {
	shapes := []shape{
		&RectangleAdapter{&Rectangle{}},
		&LineAdapter{&Line{}},
	}

	for _, singleShape := range shapes {
		singleShape.draw(10, 20, 30, 40)
	}
}
