package main

import (
	"github.com/mgaffney/tlife/life"
	"github.com/nsf/termbox-go"
)

func draw(g *life.Grid) {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			cell := g.Cell(x, y)
			switch {
			case cell == life.Alive:
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorYellow)
				// termbox.SetCell(x, y, '■', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, 'X', termbox.ColorDefault, termbox.ColorDefault)
			default:
				// termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '.', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '▁', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '┼', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '┌', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '‿', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '·', termbox.ColorDefault, termbox.ColorDefault)
				// termbox.SetCell(x, y, '·', termbox.ColorBlue|termbox.AttrBold, termbox.ColorBlack)
				termbox.SetCell(x, y, '·', termbox.ColorBlack, termbox.ColorDefault)
			}
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	g := life.New(termbox.Size())
	g.Set(1, 0, life.Alive)
	g.Set(2, 1, life.Alive)
	g.Set(0, 2, life.Alive)
	g.Set(1, 2, life.Alive)
	g.Set(2, 2, life.Alive)
	draw(g)

mainloop:
	for {
		mx, my := -1, -1
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
			if ev.Key == termbox.KeyEnter {
				g = g.Evolve()
			}
		case termbox.EventMouse:
			if ev.Key == termbox.MouseLeft {
				mx, my = ev.MouseX, ev.MouseY
				g.Toggle(mx, my)
			}
		}
		draw(g)
	}
}
