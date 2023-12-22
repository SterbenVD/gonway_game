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
	window.Run()
	frames_per_tick := config.FPS / config.TickRate
	temp := frames_per_tick
	for !window.Close() {
		if temp == 0 {
			board.Run()
			temp = frames_per_tick
		}
		temp--
	}
	rl.CloseWindow()
}
