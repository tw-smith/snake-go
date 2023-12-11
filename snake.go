package main

type SnakeSegmenter interface {
	Move()
}

type SnakeSegment struct {
	GameElement
	nextX int
	nextY int
}

func (snake *SnakeSegment) Move() {
	snake.x = snake.x + snake.nextX
	snake.y = snake.y + snake.nextY
}

func (snake *SnakeSegment) ChangeDirection(direction string) {
	switch direction {
	case "up":
		snake.nextX = 0
		snake.nextY = -1
	case "left":
		snake.nextX = -1
		snake.nextY = 0
	case "right":
		snake.nextX = 1
		snake.nextY = 0
	case "down":
		snake.nextX = 0
		snake.nextY = 1
	}
}
