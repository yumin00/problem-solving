package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	chess [][]bool
	n     int
	line  int
	cnt   int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &n)

	var nowX, nowY int
	var moveX, moveY int

	for test := 0; test < n; test++ {
		fmt.Fscanln(r, &line)
		chess = make([][]bool, line)
		for i := 0; i < line; i++ {
			chess[i] = make([]bool, line)
		}

		fmt.Fscanln(r, &nowX, &nowY)
		fmt.Fscanln(r, &moveX, &moveY)

		if nowX == moveX && nowY == moveY {
			cnt = 0
		} else {
			cnt = bfs(nowX, nowY, moveX, moveY)
		}
		fmt.Fprintln(w, cnt)
		resetChess()
	}
	w.Flush()
}

func resetChess() {
	for i := 0; i < line; i++ {
		for j := 0; j < line; j++ {
			chess[i][j] = false
		}
	}
}

type Point struct {
	x    int
	y    int
	move int
}

func bfs(nowX, nowY, moveX, moveY int) (cnt int) {
	chess[nowX][nowY] = true
	directions := [8][2]int{{-1, -2}, {-1, 2}, {-2, -1}, {-2, 1}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	queue := make([]Point, 0)
	queue = append(queue, Point{nowX, nowY, 0})
	for len(queue) > 0 {
		current := queue[0]
		for _, dir := range directions {
			dirX, dirY := current.x+dir[0], current.y+dir[1]
			if dirX == moveX && dirY == moveY {
				return current.move + 1
			} else if dirX > 0 && dirX < line && dirY > 0 && dirY < line && chess[dirX][dirY] == false {
				chess[dirX][dirY] = true
				queue = append(queue, Point{dirX, dirY, current.move + 1})
			}
		}
		queue = queue[1:]
	}
	return cnt
}
