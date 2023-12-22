package api

import rl "github.com/gen2brain/raylib-go/raylib"

// Window is a window that can be drawn to.
type Window struct {
	// Height of the window
	Height int32

	// Width of the window
	Width int32

	// Title of the window
	Title string

	// FPS of the window
	FPS int32
}

// NewWindow creates a new window.
func NewWindow(height, width, FPS int32, title string) *Window {
	return &Window{
		Height: height,
		Width:  width,
		Title:  title,
		FPS:    FPS,
	}
}

// Run runs the window.
func (w *Window) Run() {
	rl.InitWindow(w.Width, w.Height, w.Title)
	rl.SetTargetFPS(w.FPS)
}

// Close closes the window.
func (w *Window) Close() bool {
	return rl.WindowShouldClose()
}
