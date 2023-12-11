package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen   tcell.Screen
	Food     GameElement
	Snake    SnakeSegment
	Score    int
	GameOver bool
}

func (g *Game) UpdateFoodPosition(width int, height int) {
	g.Food.x = rand.Intn(width)
	g.Food.y = rand.Intn(height)
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
		g.Snake.Draw(g.Screen)
		g.Food.Draw(g.Screen)
		time.Sleep(100 * time.Millisecond)
		g.Screen.Show()
	}
}
