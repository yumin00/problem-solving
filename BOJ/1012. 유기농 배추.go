package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited [][]bool
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(r, &t)

	for i := 0; i < t; i++ {
		var m, n, k int
		fmt.Fscanln(r, &m, &n, &k)

		graph = make([][]int, n)
		for j := 0; j < n; j++ {
			graph[j] = make([]int, m)
		}
		visited = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[j] = make([]bool, m)
		}

		for l := 0; l < k; l++ {
			var x, y int
			fmt.Fscanln(r, &x, &y)
			graph[y][x] = 1
		}

		var bug int
		for j := 0; j < n; j++ {
			for l := 0; l < m; l++ {
				if graph[j][l] == 1 && !visited[j][l] {
					bug++
					dfs(j, l)
				}
			}
		}
		fmt.Fprintln(w, bug)
	}
	w.Flush()
}

func dfs(row, col int) {
	visited[row][col] = true
	if row+1 < len(graph) && graph[row+1][col] == 1 && !visited[row+1][col] {
		dfs(row+1, col)
	}
	if row-1 >= 0 && graph[row-1][col] == 1 && !visited[row-1][col] {
		dfs(row-1, col)
	}
	if col+1 < len(graph[row]) && graph[row][col+1] == 1 && !visited[row][col+1] {
		dfs(row, col+1)
	}
	if col-1 >= 0 && graph[row][col-1] == 1 && !visited[row][col-1] {
		dfs(row, col-1)
	}
}
