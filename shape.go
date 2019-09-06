package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	e      = -1
	leftP  = -2
	rightP = -3
	square = iota
	hex
	kite
)

var (
	st = math.Sqrt(3)
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

func rightShapePos(shape int, position *Position, baseSide float64) float64 {
	switch shape {
	case square:
		return position.x + baseSide
	case hex:
		return position.x + baseSide*st
	}
	return position.x
}

func downShapePos(shape int, position *Position, baseSide float64) float64 {
	switch shape {
	case square:
		return position.y + baseSide
	case hex:
		return position.y + baseSide*3/2
	}
	return position.y
}

func makeSquare(center *Position, rotation, baseSide float64, num int) *Shape {
	vertices := []*Position{
		pos(center.x-baseSide/2, center.y-baseSide/2),
		pos(center.x+baseSide/2, center.y-baseSide/2),
		pos(center.x+baseSide/2, center.y+baseSide/2),
		pos(center.x-baseSide/2, center.y+baseSide/2),
	}
	square := &Shape{vertices: vertices, center: center, rotation: rotation, num: num}
	return square
}

func makeHex(center *Position, rotation, baseSide float64, num int) *Shape {
	vertices := []*Position{
		pos(center.x, center.y-baseSide),
		pos(center.x+st*baseSide/2, center.y-baseSide/2),
		pos(center.x+st*baseSide/2, center.y+baseSide/2),
		pos(center.x, center.y+baseSide),
		pos(center.x-st*baseSide/2, center.y+baseSide/2),
		pos(center.x-st*baseSide/2, center.y-baseSide/2),
	}
	hex := &Shape{vertices: vertices, center: center, rotation: rotation, num: num}
	return hex
}
