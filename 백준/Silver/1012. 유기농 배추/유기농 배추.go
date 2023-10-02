package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	mainGraph   [][]int
	mainVisited [][]bool
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(r, &t)

	for i := 0; i < t; i++ {
		var m, n, k int
		fmt.Fscanln(r, &m, &n, &k)

		mainGraph = make([][]int, n)
		for j := 0; j < n; j++ {
			mainGraph[j] = make([]int, m)
		}
		mainVisited = make([][]bool, n)
		for j := 0; j < n; j++ {
			mainVisited[j] = make([]bool, m)
		}

		for l := 0; l < k; l++ {
			var x, y int
			fmt.Fscanln(r, &x, &y)
			mainGraph[y][x] = 1
		}

		var numOfWorm int
		for j := 0; j < n; j++ {
			for l := 0; l < m; l++ {
				if mainGraph[j][l] == 1 && !mainVisited[j][l] {
					numOfWorm++
					mainDFS(j, l)
				}
			}
		}
		fmt.Fprintln(w, numOfWorm)
	}
	w.Flush()
}

func mainDFS(row, col int) {
	mainVisited[row][col] = true
	if row+1 < len(mainGraph) && mainGraph[row+1][col] == 1 && !mainVisited[row+1][col] {
		mainDFS(row+1, col)
	}
	if row-1 >= 0 && mainGraph[row-1][col] == 1 && !mainVisited[row-1][col] {
		mainDFS(row-1, col)
	}
	if col+1 < len(mainGraph[row]) && mainGraph[row][col+1] == 1 && !mainVisited[row][col+1] {
		mainDFS(row, col+1)
	}
	if col-1 >= 0 && mainGraph[row][col-1] == 1 && !mainVisited[row][col-1] {
		mainDFS(row, col-1)
	}
}
