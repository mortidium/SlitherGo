package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	numWidth  = 18
	numHeight = 40
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

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func screenshot(key string) {
	scrot := exec.Cmd{
		Path: "/usr/bin/scrot",
		Args: []string{
			"'ok.png'",
			"--focused",
		},
	}
	err := scrot.Run()
	failOnErr(err, "Failed to make screenshot")
	currDir, err := os.Getwd()
	failOnErr(err, "couldn't find curr dir")
	files, err := filePathWalkDir(currDir)
	failOnErr(err, "couldn't get files")
	ss := ""
	for _, f := range files {
		if strings.HasSuffix(f, ".png") {
			ss = f
			break
		}
	}
	if ss == "" {
		failOnErr(fmt.Errorf("Kurwa"), "O nie")
	}

	time.Sleep(100 * time.Millisecond)
	err = os.Rename(ss, fmt.Sprintf("%v/images/%v.png", currDir, key))
	failOnErr(err, "renaming")
}

func main() {
	key := "SqSmHa1"
	if len(os.Args) > 1 {
		key = os.Args[1]
	}

	examples := loadExamples()
	config, ok := examples[key]
	if !ok {
		failOnErr(fmt.Errorf(key), "Example not found")
	}

	err := sdl.Init(sdl.INIT_EVERYTHING)
	failOnErr(err, "Failed initializing sdl")
	defer sdl.Quit()

	err = ttf.Init()
	failOnErr(err, "Failed font init")
	defer ttf.Quit()

	window, err := sdl.CreateWindow(
		"Slitherlink",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(config.width), int32(config.height),
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

	drawBoard(renderer, config)

	renderer.Present()

	window.Raise()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Not saved")
				return
			}
		}
		ks := sdl.GetKeyboardState()
		if ks[sdl.SCANCODE_RETURN] == 1 {
			fmt.Printf("Saved: %v\n", key)
			screenshot(key)
			return
		}
	}
}
