package main

import (
	"github.com/gdamore/tcell"
)

type GameElementer interface {
	Draw()
	CheckCollision()
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

func (element GameElement) CheckCollision(targetList []GameElement) bool {
	for _, target := range targetList {
		if element.x == target.x && element.y == target.y {
			return true
		}
	}
	return false
}

type Food struct {
	GameElement
}
