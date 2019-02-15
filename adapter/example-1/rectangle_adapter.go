package main

type RectangleAdapter struct {
	rect shape
}

func (l *RectangleAdapter) draw(x, y, z, j int) {
	l.rect.draw(x, y, z, j)
}
