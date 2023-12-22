package main

import (
	"github.com/SterbenVD/gonway_game/api"
	"github.com/SterbenVD/gonway_game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	height := config.Height
	width := config.Width
	title := config.Title
	FPS := config.FPS
	window := api.NewWindow(height, width, FPS, title)
	board := api.NewBoard(height, width, 10, rl.Black, rl.White)
	board.Init()
	window.Run()
	for !window.Close() {
		if rl.IsKeyPressed(rl.KeySpace) {
			board.Init()
		}
		if rl.IsKeyPressedRepeat(rl.KeyRight) {
			window.ChangeFPS(true)
		}
		if rl.IsKeyPressedRepeat(rl.KeyLeft) {
			window.ChangeFPS(false)
		}
		board.Run()
	}
	rl.CloseWindow()
}
