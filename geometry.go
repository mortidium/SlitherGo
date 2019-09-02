package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Position gives a point on plane
type Position struct {
	x, y float64
}

func pos(x, y float64) *Position { return &Position{x: x, y: y} }
func move(p, q *Position) *Position {
	return &Position{x: p.x + q.x, y: p.y + q.y}
}

func drawLine(renderer *sdl.Renderer, from, to *Position) {
	setColor("line", renderer)
	renderer.DrawLine(int32(from.x), int32(from.y), int32(to.x), int32(to.y))
}
