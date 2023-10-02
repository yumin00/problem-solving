package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n       int
	graph   [][]int
	visited []bool
	N, M, V int
	r       *bufio.Reader
	w       *bufio.Writer
)

func main() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &N, &M, &V)

	graph = make([][]int, N+1)
	visited = make([]bool, N+1)

	for i := range graph {
		graph[i] = make([]int, N+1)
	}

	//graph 입력받기, check false로 만들기
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		graph[x][y] = 1
		graph[y][x] = 1
	}

	dfs(V)
	w.Flush()
	resetVisited()
	fmt.Println()

	bfs(V)
}

func dfs(V int) {
	visited[V] = true
	fmt.Fprint(w, V, " ")
	for i := 0; i < len(graph[V]); i++ {
		if graph[V][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}

func bfs(V int) {
	visited[V] = true
	q := []int{V}

	for len(q) != 0 {
		front := q[0]
		fmt.Print(front, " ")
		q = q[1:]
		for i := 0; i <= N; i++ {
			if graph[front][i] == 1 && !visited[i] {
				visited[i] = true
				q = append(q, i)
			}
		}
	}
}

func resetVisited() {
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
}
