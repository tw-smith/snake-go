package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen   tcell.Screen
	Food     GameElement
	Snake    Snake
	Score    int
	GameOver bool
}

func (g *Game) UpdateFoodPosition(width int, height int) {
	g.Food.x = rand.Intn(width)
	g.Food.y = rand.Intn(height)
}

func (g *Game) DrawBorders(s tcell.Screen, width int, height int) {
	borderStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorPurple)
	for i := 0; i < height; i++ {
		s.SetContent(0, i, tcell.RuneVLine, nil, borderStyle)
	}
	for i := 0; i < width; i++ {
		s.SetContent(i, 0, tcell.RuneHLine, nil, borderStyle)
	}
	for i := 0; i < height; i++ {
		s.SetContent(width, i, tcell.RuneVLine, nil, borderStyle)
	}
	for i := 0; i < width; i++ {
		s.SetContent(i, height, tcell.RuneHLine, nil, borderStyle)
	}

}

func (g *Game) DrawScore(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, char := range text {
		s.SetContent(col, row, char, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) DrawSnake(s tcell.Screen, snake Snake) {
	for _, segment := range snake.segments {
		segment.Draw(s)
	}
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	g.UpdateFoodPosition(width, height)
	g.GameOver = false
	g.Score = 0

	for {
		g.Screen.Clear()
		g.Snake.Move()
		g.DrawBorders(g.Screen, width, height)
		g.DrawSnake(g.Screen, g.Snake)
		g.Food.Draw(g.Screen)
		if g.Snake.segments[0].CheckCollision([]GameElement{g.Food}) {
			g.Score++
			g.UpdateFoodPosition(width, height)
			g.Snake.AddSegment()
		}
		g.DrawScore(g.Screen, 7, 2, 20, 2, fmt.Sprintf("Score: %d", g.Score))
		time.Sleep(100 * time.Millisecond)
		g.Screen.Show()
	}
}
