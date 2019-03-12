package life

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		w, h int
	}{
		{0, 0},
		{10, 10},
		{10, 100},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%dx%d", tt.w, tt.h), func(t *testing.T) {
			g := New(tt.w, tt.h)
			gw, gh := g.Size()
			if gw != tt.w || gh != tt.h {
				t.Errorf("expected w: %d h: %d, got w: %d h:%d", tt.w, tt.h, gw, gw)
			}
			for y := 0; y < tt.h; y++ {
				for x := 0; x < tt.w; x++ {
					c := g.Cell(x, y)
					if c != Dead {
						t.Errorf("cell(%d,%d): expected %v, got %v", x, y, Dead, c)
					}
				}
			}
		})
	}
}

func TestResize(t *testing.T) {
	type gs struct {
		w, h int
	}
	var tests = []struct {
		a, b gs
	}{
		{gs{0, 0}, gs{10, 10}},
		{gs{10, 10}, gs{0, 0}},
		{gs{5, 5}, gs{10, 10}},
		{gs{10, 10}, gs{5, 5}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%dx%d=>%dx%d", tt.a.w, tt.a.h, tt.b.w, tt.b.h), func(t *testing.T) {
			ga := New(tt.a.w, tt.a.h)
			for y := 0; y < tt.a.h; y++ {
				for x := 0; x < tt.a.w; x++ {
					ga.Toggle(x, y)
				}
			}
			t.Log(ga)
			gb := ga.Resize(tt.b.w, tt.b.h)

			gbw, gbh := gb.Size()
			if gbw != tt.b.w || gbh != tt.b.h {
				t.Errorf("expected w: %d h: %d, got w: %d h:%d", tt.b.w, tt.b.h, gbw, gbw)
			}
			t.Log(gb)
			for y := 0; y < tt.b.h; y++ {
				for x := 0; x < tt.b.w; x++ {
					if got, want := gb.Cell(x, y), ga.Cell(x, y); got != want {
						t.Errorf("cell(%d,%d): expected %v, got %v", x, y, want, got)
					}
				}
			}
		})
	}
}

func TestSetCell(t *testing.T) {
	w, h := 10, 100
	g := New(w, h)

	var c Cell
	x, y := 6, 9
	c = g.Cell(x, y)
	if c != Dead {
		t.Errorf("default cell(%d,%d): expected %v, got %v", x, y, Dead, c)
	}
	g.Set(x, y, Alive)
	c = g.Cell(x, y)
	if c != Alive {
		t.Errorf("set cell(%d,%d): expected %v, got %v", x, y, Alive, c)
	}

	g.Toggle(x, y)
	c = g.Cell(x, y)
	if c != Dead {
		t.Errorf("toggle cell(%d,%d): expected %v, got %v", x, y, Dead, c)
	}

}

func TestOutOfBoundsSetCell(t *testing.T) {
	w, h := 10, 10
	g := New(w, h)

	var tests = []struct {
		x, y int
	}{
		{-1, 0},
		{0, -1},
		{-1, -1},
		{11, 10},
		{10, 11},
		{11, 11},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("x:%d,y:%d", tt.x, tt.y), func(t *testing.T) {
			var c Cell
			c = g.Cell(tt.x, tt.y)
			if c != Dead {
				t.Errorf("default value, cell(%d,%d): expected %v, got %v", tt.x, tt.y, Dead, c)
			}
			g.Set(tt.x, tt.y, Alive)
			c = g.Cell(tt.x, tt.y)
			if c != Dead {
				t.Errorf("set value, cell(%d,%d): expected %v, got %v", tt.x, tt.y, Dead, c)
			}
			g.Toggle(tt.x, tt.y)
			c = g.Cell(tt.x, tt.y)
			if c != Dead {
				t.Errorf("toggle value, cell(%d,%d): expected %v, got %v", tt.x, tt.y, Dead, c)
			}
		})
	}
}

type cc struct {
	x, y int
}

func TestEvolve(t *testing.T) {
	w, h := 3, 3
	var tests = []struct {
		name   string
		before []cc
		after  []cc
	}{
		{"empty", []cc{}, []cc{}},
		{"one-cell", []cc{cc{1, 1}}, []cc{}},
		{"two-cell",
			[]cc{
				cc{1, 1},
				cc{1, 2},
			},
			[]cc{}},
		{"three-cell",
			[]cc{
				cc{1, 1},
				cc{1, 2},
				cc{2, 1},
			},
			[]cc{
				cc{1, 1},
				cc{1, 2},
				cc{2, 1},
				cc{2, 2},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gb := New(w, h)
			for _, c := range tt.before {
				gb.Set(c.x, c.y, Alive)
			}
			ga := New(w, h)
			for _, c := range tt.after {
				ga.Set(c.x, c.y, Alive)
			}
			actual := gb.Evolve()
			if !reflect.DeepEqual(actual, ga) {
				t.Errorf("(%v): expected %v, actual %v", gb, ga, actual)
			}

		})
	}
}
