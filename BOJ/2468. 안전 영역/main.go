package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var visited [][]bool
var graph [][]int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscan(r, &n)

	graph = make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(r, &a)
			graph[i][j] = a
		}
	}

	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	var cnt int

	maxHeight := getMaxHeight(graph)
	maxSafeArea := 0

	for h := 0; h <= maxHeight; h++ {
		checkSafeArea(h)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if visited[i][j] == false {
					bfs(i, j)
					cnt++
				}
			}
		}
		if cnt > maxSafeArea {
			maxSafeArea = cnt
		}
		resetVisited()
		cnt = 0
	}
	fmt.Fprintln(w, maxSafeArea)
	w.Flush()
}

type Point struct {
	row    int
	column int
}

func resetVisited() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}
}

func bfs(startRow int, startColumn int) {
	visited[startRow][startColumn] = true
	//w := bufio.NewWriter(os.Stdout)
	queue := make([]Point, 0)
	queue = append(queue, Point{startRow, startColumn})

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		for _, dir := range directions {
			x, y := current.row+dir[0], current.column+dir[1]
			if x >= 0 && x < n && y >= 0 && y < n && visited[x][y] == false {
				visited[x][y] = true
				queue = append(queue, Point{x, y})
			}
		}
		queue = queue[1:]
	}
}

func checkSafeArea(height int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] <= height {
				visited[i][j] = true
			}
		}
	}
}

func getMaxHeight(graph [][]int) int {
	var maxHeight int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				maxHeight = graph[i][j]
			} else if graph[i][j] > maxHeight {
				maxHeight = graph[i][j]
			}
		}
	}
	return maxHeight
}
