package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	e      = -1
	square = iota
	hex
	kite
)

// Shape crates field shape
type Shape struct {
	vertices []*Position
	center   *Position
	rotation float64
	num      int
	kind     int
}

func (s *Shape) draw(renderer *sdl.Renderer) {
	if len(s.vertices) == 0 {
		return
	}
	// draw lines
	initPos := s.vertices[0]
	lastPos := s.vertices[0]
	for _, p := range s.vertices {
		drawLine(renderer, lastPos, p)
		lastPos = p
	}
	drawLine(renderer, lastPos, initPos)
	// draw vertices
	for _, p := range s.vertices {
		d, err := initDot(renderer, pos(p.x, p.y))
		failOnErr(err, "unable to render dot")
		d.onDraw(renderer)
	}
	// draw number
	if s.num != e {
		printNum(renderer, s.num, s.center)
	}
}

func nextShapePos(shape int, pos *Position, baseSide float64) *Position {
	switch shape {
	case square:
		return &Position{x: pos.x + baseSide, y: pos.y}
	}
	return pos
}

func downShapePos(shape int, pos *Position, baseSide float64) *Position {
	switch shape {
	case square:
		return &Position{x: pos.x, y: pos.y + baseSide}
	}
	return pos
}

func (s *Shape) downLShapePos(baseSide float64) *Position {
	switch s.kind {
	case square:
		return &Position{x: s.center.x, y: s.center.y + baseSide}
	}
	return s.center
}

func (s *Shape) downRShapePos(baseSide float64) *Position {
	switch s.kind {
	case square:
		return &Position{x: s.center.x, y: s.center.y + baseSide}
	}
	return s.center
}

func makeSquare(center *Position, rotation, baseSide float64, num int) *Shape {
	vertices := []*Position{
		&Position{x: center.x - baseSide/2, y: center.y - baseSide/2},
		&Position{x: center.x + baseSide/2, y: center.y - baseSide/2},
		&Position{x: center.x + baseSide/2, y: center.y + baseSide/2},
		&Position{x: center.x - baseSide/2, y: center.y + baseSide/2},
	}
	square := &Shape{vertices: vertices, center: center, rotation: rotation, num: num}
	return square
}
