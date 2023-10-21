package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n       int
	game    [][]int
	visited [][]bool
	check   bool
)

func main() {
	check = false
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &n)

	game = make([][]int, n)
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		game[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(r, &a)
			game[i][j] = a
		}
		fmt.Fscanln(r)
	}

	dfs(0, 0)

	if check == true {
		fmt.Fprintln(w, "HaruHaru")
	} else {
		fmt.Fprintln(w, "Hing")
	}

	w.Flush()

}

func dfs(x, y int) {
	if game[x][y] == -1 {
		check = true
		return
	}
	visited[x][y] = true
	plus := game[x][y]

	if plus == 0 {
		return
	}

	if y+plus < n && visited[x][y+plus] == false {
		dfs(x, y+plus)
	}
	if x+plus < n && visited[x+plus][y] == false {
		dfs(x+plus, y)
	}
}
