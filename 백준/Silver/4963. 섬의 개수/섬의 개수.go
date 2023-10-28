package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	width, height int
	box           [][]bool
	cnt           int
	result        []int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for {
		cnt = 0
		fmt.Fscanln(r, &width, &height)
		if width == 0 && height == 0 {
			break
		}

		box = make([][]bool, height)

		for i := 0; i < height; i++ {
			box[i] = make([]bool, width)
			for j := 0; j < width; j++ {
				var a int
				fmt.Fscan(r, &a)
				if a == 0 {
					box[i][j] = false
				} else {
					box[i][j] = true
				}
			}
			fmt.Fscanln(r)
		}

		count()
		result = append(result, cnt)
	}

	for i := 0; i < len(result); i++ {
		fmt.Fprintln(w, result[i])
	}

	w.Flush()
}

func dfs(x, y int) {
	directiosn := [8][2]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}
	box[x][y] = false

	for _, v := range directiosn {
		moveX := x + v[0]
		moveY := y + v[1]
		if moveX < height && moveX >= 0 && moveY < width && moveY >= 0 && box[moveX][moveY] == true {
			dfs(moveX, moveY)
		}

	}
}

func count() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if box[i][j] == true {
				dfs(i, j)
				cnt++
			}
		}
	}
}
