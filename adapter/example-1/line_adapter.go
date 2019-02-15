package main

type LineAdapter struct {
	line shape
}

func (l *LineAdapter) draw(x, y, z, j int) {
	l.line.draw(x, y, z, j)
}
