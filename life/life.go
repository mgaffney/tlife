// Package life implements Conway's Game of Life.
package life

import "bytes"

//go:generate stringer -type=Cell

// A Cell represents a cell in Conway's Game of Life. It can be either
// alive or dead.
type Cell int

const (
	Dead Cell = iota
	Alive
)

// Grid is a finite rectangular grid of Cells.
type Grid struct {
	w, h int
	g    [][]Cell
}

// New creates a new Grid of width and height.
func New(width, height int) *Grid {
	grid := make([][]Cell, height)
	cells := make([]Cell, width*height)
	for i := range grid {
		grid[i], cells = cells[:width], cells[width:]
	}
	return &Grid{
		w: width,
		h: height,
		g: grid,
	}
}

// Size returns the width and height of g.
func (g *Grid) Size() (width, height int) {
	return g.w, g.h
}

// Resize returns a new grid of width and height with an initial state
// based on g.
func (g *Grid) Resize(width, height int) *Grid {
	b := New(width, height)
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			b.Set(x, y, g.Cell(x, y))
		}
	}
	return b
}

// Cell returns the cell at x,y.
func (g *Grid) Cell(x, y int) Cell {
	if x < 0 || x >= g.w {
		return Dead
	}
	if y < 0 || y >= g.h {
		return Dead
	}
	return g.g[y][x]
}

// Set sets the cell at x,y to c.
func (g *Grid) Set(x, y int, c Cell) {
	if x < 0 || x >= g.w {
		return
	}
	if y < 0 || y >= g.h {
		return
	}
	g.g[y][x] = c
}

// Toggle toggles the value of the cell at x,y.
func (g *Grid) Toggle(x, y int) {
	c := g.Cell(x, y)
	switch c {
	case Dead:
		g.Set(x, y, Alive)
	case Alive:
		g.Set(x, y, Dead)
	}

}

// Evolve returns a grid corresponding to Conway's "genetic laws" applied
// to g.
//
// Evolve iterates over each cell in the grid applying the following
// transitions:
//
// 1. Any live cell with fewer than two live neighbors dies, as if by
// underpopulation.
//
// 2. Any live cell with two or three live neighbors lives on to the next
// generation.
//
// 3. Any live cell with more than three live neighbors dies, as if by
// overpopulation.
//
// 4. Any dead cell with exactly three live neighbors becomes a live cell,
// as if by reproduction.
func (g *Grid) Evolve() *Grid {
	r := New(g.Size())

	for y, yv := range g.g {
		for x, xv := range yv {
			live := g.liveNeighbors(x, y)
			if xv == Dead && live == 3 {
				r.Set(x, y, Alive)
				continue
			}
			if xv == Alive && live == 2 || live == 3 {
				r.Set(x, y, Alive)
				continue
			}
		}
	}
	return r
}

func (g *Grid) liveNeighbors(x, y int) int {
	var live int
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.Cell(x+dx, y+dy) == Alive {
				live++
			}
		}
	}
	return live
}

func (g *Grid) String() string {
	var buf bytes.Buffer
	for _, yv := range g.g {
		buf.WriteString("\n")
		for _, xv := range yv {
			switch xv {
			case Alive:
				buf.WriteString("X")
			default:
				buf.WriteString(".")
			}
		}
	}
	buf.WriteString("\n")
	return buf.String()
}
