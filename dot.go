package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	texWidth  = 21
	texHeight = 21
)

// Dot draws a dot on position
type Dot struct {
	pos *Position
	tex *sdl.Texture
}

func (d *Dot) onDraw(renderer *sdl.Renderer) {
	renderer.Copy(
		d.tex,
		&sdl.Rect{X: 0, Y: 0, W: texWidth, H: texHeight},
		&sdl.Rect{X: int32(d.pos.x - texWidth/2), Y: int32(d.pos.y - texHeight/2), W: texWidth, H: texHeight})
}

func initDot(renderer *sdl.Renderer, position *Position) (*Dot, error) {
	img, err := sdl.LoadBMP("sprites/circFull.bmp")
	if err != nil {
		return nil, fmt.Errorf("LoadBMP(%v): %v", "circFull", err)
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("CreateTexture(%v): %v", "dot", err)
	}
	dot := &Dot{pos: position, tex: texture}
	return dot, nil
}
