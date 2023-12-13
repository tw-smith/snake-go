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
	// -1 is here to stop food spawning right on the border and therefore being unreachable
	g.Food.x = rand.Intn(width - 1)
	g.Food.y = rand.Intn(height - 1)
}

func (g *Game) DrawBorders(s tcell.Screen, gameAreaWidth int, gameAreaHeight int) {
	borderStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorPurple)
	for i := 0; i < gameAreaHeight; i++ {
		s.SetContent(0, i, tcell.RuneVLine, nil, borderStyle)
	}
	for i := 0; i < gameAreaWidth; i++ {
		s.SetContent(i, 0, tcell.RuneHLine, nil, borderStyle)
	}
	for i := 0; i < gameAreaHeight; i++ {
		s.SetContent(gameAreaWidth, i, tcell.RuneVLine, nil, borderStyle)
	}
	for i := 0; i < gameAreaWidth; i++ {
		s.SetContent(i, gameAreaHeight, tcell.RuneHLine, nil, borderStyle)
	}
	s.SetContent(0, 0, tcell.RuneULCorner, nil, borderStyle)
	s.SetContent(0, gameAreaHeight, tcell.RuneLLCorner, nil, borderStyle)
	s.SetContent(gameAreaWidth, gameAreaHeight, tcell.RuneLRCorner, nil, borderStyle)
	s.SetContent(gameAreaWidth, 0, tcell.RuneURCorner, nil, borderStyle)

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

func (g *Game) DrawGameOver(s tcell.Screen, gameAreaHeight, gameAreaWidth, score int) {
	row := gameAreaHeight / 2
	col := gameAreaWidth / 2
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	message := "Game Over!"
	for _, char := range message {
		s.SetContent(col, row, char, nil, style)
		col++
	}
	s.Show()
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	gameAreaWidth := width / 2
	gameAreaHeight := height / 2
	g.UpdateFoodPosition(gameAreaWidth, gameAreaHeight)
	g.GameOver = false
	g.Score = 0

	for {
		g.Screen.Clear()

		g.DrawBorders(g.Screen, gameAreaWidth, gameAreaHeight)
		g.DrawSnake(g.Screen, g.Snake)
		g.Food.Draw(g.Screen)

		snakeHead := g.Snake.segments[len(g.Snake.segments)-1]

		if snakeHead.CheckCollision([]GameElement{g.Food}) {
			g.Score++
			g.UpdateFoodPosition(gameAreaWidth, gameAreaHeight)
			g.Snake.AddSegment()
		}

		if snakeHead.x >= gameAreaWidth || snakeHead.x <= 0 || snakeHead.y >= gameAreaHeight || snakeHead.y <= 0 {
			g.DrawGameOver(g.Screen, gameAreaHeight, gameAreaWidth, g.Score)
			break
		}

		if snakeHead.CheckCollision(g.Snake.segments[:(len(g.Snake.segments) - 2)]) {
			g.DrawGameOver(g.Screen, gameAreaHeight, gameAreaWidth, g.Score)
			break
		}
		g.Snake.Move()
		g.DrawScore(g.Screen, 7, 2, 20, 2, fmt.Sprintf("Score: %d", g.Score))
		time.Sleep(100 * time.Millisecond)
		g.Screen.Show()
	}
}
