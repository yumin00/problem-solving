package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n    int
	m    int
	maze [][]bool
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &n, &m)

	maze = make([][]bool, n)
	for i := 0; i < n; i++ {
		maze[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscanf(r, "%1d", &a)
			if a == 1 {
				maze[i][j] = true
			}
		}
		fmt.Fscanln(r)
	}

	result := bfs(0, 0)
	fmt.Fprintln(w, result)
	w.Flush()
}

type Point struct {
	row      int
	column   int
	distance int
}

func bfs(i, j int) (result int) {
	maze[i][j] = false
	var q []Point
	q = append(q, Point{
		row:      i,
		column:   j,
		distance: 1,
	})

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		currentX := q[0].row
		currentY := q[0].column
		currentDistance := q[0].distance
		for _, v := range directions {
			moveX := currentX + v[0]
			moveY := currentY + v[1]
			if moveX == n-1 && moveY == m-1 {
				return currentDistance + 1
			}
			if moveX >= 0 && moveX < n && moveY >= 0 && moveY < m && maze[moveX][moveY] == true {
				maze[moveX][moveY] = false
				q = append(q, Point{
					row:      moveX,
					column:   moveY,
					distance: currentDistance + 1,
				})
			}
		}
		q = q[1:]
	}
	return
}
