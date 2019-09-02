package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth  = 1000
	screenHeight = 1000
	numWidth     = 18
	numHeight    = 40
)

var (
	font *ttf.Font
)

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, ":", err)
		os.Exit(1)
	}
}
func resetColor(r *sdl.Renderer) {
	r.SetDrawColor(255, 255, 255, 255)
}

func setColor(color string, r *sdl.Renderer) {
	switch strings.ToLower(color) {
	case "dot":
		r.SetDrawColor(100, 100, 100, 255)
	case "line":
		r.SetDrawColor(170, 170, 170, 255)
	default:
		resetColor(r)
	}
}

func printNum(renderer *sdl.Renderer, num int, pos *Position) {
	texture := createText(fmt.Sprintf("%d", num), renderer)

	renderer.Copy(texture, nil, &sdl.Rect{X: int32(pos.x - numWidth/2), Y: int32(pos.y - numHeight/2), W: numWidth, H: numHeight})
}

func createText(text string, renderer *sdl.Renderer) *sdl.Texture {
	solidSurface, err := font.RenderUTF8Blended(text, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	failOnErr(err, "Failed to render text")

	solidTexture, err := renderer.CreateTextureFromSurface(solidSurface)
	failOnErr(err, "Failed to create texture")
	return solidTexture
}

func drawBoard(renderer *sdl.Renderer) {
	board := makeSquareBoard(pos(100, 100), 80, SqSmHa1)
	board.draw(renderer)
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	failOnErr(err, "Failed initializing sdl")
	defer sdl.Quit()

	err = ttf.Init()
	failOnErr(err, "Failed font init")
	defer ttf.Quit()

	window, err := sdl.CreateWindow(
		"Slitherlink",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL,
	)
	failOnErr(err, "Failed creating window")
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	failOnErr(err, "Failed creating renderer")
	defer renderer.Destroy()

	font, err = ttf.OpenFont("sprites/dejaVuSansBold.ttf", 40)
	failOnErr(err, "Failed to open font")
	defer font.Close()

	resetColor(renderer)
	renderer.Clear()

	drawBoard(renderer)

	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
