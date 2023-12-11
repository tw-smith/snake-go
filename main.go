package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

// const (
// 	screenWidth  = 320
// 	screenHeight = 240
// 	tilesize     = 5
// )

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snake := SnakeSegment{
		GameElement{
			x:     7,
			y:     7,
			sizex: 5,
			sizey: 5,
			style: tcell.StyleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorYellow),
		},
		1,
		0,
	}

	game := Game{
		Screen:   screen,
		Snake:    snake,
		GameOver: false,
		Score:    0,
	}

	go game.Run()
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize: //TODO check what this does
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			}
			if event.Key() == tcell.KeyUp {
				game.Snake.ChangeDirection("up")
			}
			if event.Key() == tcell.KeyLeft {
				game.Snake.ChangeDirection("left")
			}
			if event.Key() == tcell.KeyRight {
				game.Snake.ChangeDirection("right")
			}
			if event.Key() == tcell.KeyDown {
				game.Snake.ChangeDirection("down")
			}
		}
	}

	// testEl := Food{
	// 	GameElement{
	// 		x:     5,
	// 		y:     5,
	// 		sizex: 5,
	// 		sizey: 5,
	// 		style: tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorBlue),
	// 	},
	// }

	// testEl.Draw(screen)
	// screen.Show()

	// snakeBody := SnakeBody{
	// 	x:      5,
	// 	y:      10,
	// 	xSpeed: 1,
	// 	ySpeed: 0,
	// }

}
