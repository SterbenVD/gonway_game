package api

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Board is a conway's game of life board.
type Board struct {
	// Height of the board
	Height int32

	// Width of the board
	Width int32

	// Cells of the board
	Cells [][]bool

	// CellSize of the board
	CellSize int32

	// Number of cells in the board per row
	CellsPerRow int32

	// Number of cells in the board per column
	CellsPerColumn int32

	// CellColor of the board
	CellColor rl.Color

	// BackgroundColor of the board
	BackgroundColor rl.Color
}

// NewBoard creates a new board.
func NewBoard(height, width, cellSize int32, cellColor, backgroundColor rl.Color) *Board {
	_cellsperrow := width / cellSize
	_cellspercolumn := height / cellSize

	Cells := make([][]bool, _cellspercolumn)
	for i := range Cells {
		Cells[i] = make([]bool, _cellsperrow)
	}

	return &Board{
		Height:          height,
		Width:           width,
		CellSize:        cellSize,
		CellsPerRow:     _cellsperrow,
		CellsPerColumn:  _cellspercolumn,
		Cells:           Cells,
		CellColor:       cellColor,
		BackgroundColor: backgroundColor,
	}
}

func (b *Board) Init() {
	for i := range b.Cells {
		for j := range b.Cells[i] {
			b.Cells[i][j] = rand.Intn(2) == 1
		}
	}

}

// Run runs the board.
func (b *Board) Run() {
	b.UpdateCells()
	rl.BeginDrawing()
	rl.ClearBackground(b.BackgroundColor)
	b.drawCells()
	rl.EndDrawing()
}

// DrawCells draws the cells of the board.
func (b *Board) drawCells() {
	for i := int32(0); i < b.CellsPerColumn; i++ {
		for j := int32(0); j < b.CellsPerRow; j++ {
			if b.Cells[i][j] {
				rl.DrawRectangle(j*b.CellSize, i*b.CellSize, b.CellSize, b.CellSize, b.CellColor)
			}
		}
	}
}

// UpdateCells updates the cells of the board.
func (b *Board) UpdateCells() {
	new_b := NewBoard(b.Height, b.Width, b.CellSize, b.CellColor, b.BackgroundColor)
	for i := int32(0); i < b.CellsPerColumn; i++ {
		for j := int32(0); j < b.CellsPerRow; j++ {
			neigh := b.getNeighbours(i, j)
			if b.Cells[i][j] {
				if neigh == 2 || neigh == 3 {
					new_b.Cells[i][j] = true
				} else {
					new_b.Cells[i][j] = false
				}
			} else {
				if neigh == 3 {
					new_b.Cells[i][j] = true
				} else {
					new_b.Cells[i][j] = false
				}
			}
		}
	}
	b.Cells = new_b.Cells
}

// GetNeighbours gets the neighbours of a cell.
func (b *Board) getNeighbours(i, j int32) int32 {
	neigh := int32(0)
	if i != 0 && b.Cells[i-1][j] {
		neigh++
	}
	if j != 0 && b.Cells[i][j-1] {
		neigh++
	}
	if i != 0 && j != 0 && b.Cells[i-1][j-1] {
		neigh++
	}
	if i != b.CellsPerColumn-1 && b.Cells[i+1][j] {
		neigh++
	}
	if j != b.CellsPerRow-1 && b.Cells[i][j+1] {
		neigh++
	}
	if i != b.CellsPerColumn-1 && j != b.CellsPerRow-1 && b.Cells[i+1][j+1] {
		neigh++
	}
	if i != 0 && j != b.CellsPerRow-1 && b.Cells[i-1][j+1] {
		neigh++
	}
	if i != b.CellsPerColumn-1 && j != 0 && b.Cells[i+1][j-1] {
		neigh++
	}
	return neigh
}
