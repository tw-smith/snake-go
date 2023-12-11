package main

import (
	"github.com/gdamore/tcell"
)

// type Game struct {
// 	Screen   tcell.Screen
// 	Food     GameElement
// 	Snake	 SnakeSegment
// 	Score    int
// 	GameOver bool
// }

type GameElementer interface {
	Draw()
}

type SnakeSegmenter interface {
	Move()
}

type GameElement struct {
	x     int
	y     int
	sizex int
	sizey int
	style tcell.Style
}

func (x *GameElement) Draw(s tcell.Screen) {
	s.SetContent(x.x, x.y, '\u25CF', nil, x.style)
}

func (x *SnakeSegment) Move() {
	x.x = x.x + 10
}

type Food struct {
	GameElement
}

type SnakeSegment struct {
	GameElement
}

// func (g *Game) Run() {
// 	width, height := g.Screen.Size()
// 	g.UpdateFoodPosition(width, height)
// 	g.GameOver = false
// 	g.Score = 0
// 	g.Food.Draw(g.Screen)
// 	g.Screen.Show()
// }
