package main

type SnakeSegmenter interface {
	Move()
}

type SnakeSegment struct {
	GameElement
}

type Snake struct {
	segments  []SnakeSegment
	xVelocity int
	yVelocity int
}

func (snake *Snake) Move() {
	to_append := SnakeSegment{
		GameElement{
			x: snake.segments[len(snake.segments)-1].x + snake.xVelocity,
			y: snake.segments[len(snake.segments)-1].y + snake.yVelocity,
		},
	}
	snake.segments = append(snake.segments, to_append)
	snake.segments = snake.segments[1:]
}

func (snake *Snake) ChangeDirection(direction string) {
	switch direction {
	case "up":
		snake.xVelocity = 0
		snake.yVelocity = -1
	case "left":
		snake.xVelocity = -1
		snake.yVelocity = 0
	case "right":
		snake.xVelocity = 1
		snake.yVelocity = 0
	case "down":
		snake.xVelocity = 0
		snake.yVelocity = 1
	}
}

func (snake *Snake) AddSegment() {
	to_prepend := SnakeSegment{
		GameElement{
			x: snake.segments[0].x,
			y: snake.segments[0].y,
		},
	}

	snake.segments = append([]SnakeSegment{to_prepend}, snake.segments...)
}
