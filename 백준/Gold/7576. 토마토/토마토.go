package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y, max int
}

var (
	n, m int
	box  [][]int
	q    []Point
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscan(r, &n, &m)

	box = make([][]int, m)
	for i := range box {
		box[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &box[i][j])
			if box[i][j] == 1 {
				q = append(q, Point{x: i, y: j})
			}
		}
	}

	max := bfs()
	cnt := checkTomato(max)

	fmt.Fprintln(w, cnt)
	w.Flush()
}

func bfs() int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var max int

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		for _, dir := range directions {
			moveX, moveY := current.x+dir[0], current.y+dir[1]
			if 0 <= moveX && moveX < m && 0 <= moveY && moveY < n && box[moveX][moveY] == 0 {
				box[moveX][moveY] = 1
				q = append(q, Point{x: moveX, y: moveY, max: current.max + 1})
			}
		}
		max = current.max
	}
	return max
}

func checkTomato(cnt int) int {
	for _, row := range box {
		for _, val := range row {
			if val == 0 {
				return -1
			}
		}
	}
	return cnt
}
