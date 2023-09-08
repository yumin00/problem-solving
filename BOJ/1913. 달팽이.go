package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	var location int
	var locationX int
	var locationY int

	fmt.Fscan(r, &n)
	fmt.Fscan(r, &location)

	snail := make([][]int, n)

	x := 0
	y := 0
	idx := 0
	dx := make([]int, 4)
	dx[0] = 1  //아래
	dx[1] = 0  //오른쪽
	dx[2] = -1 // 위
	dx[3] = 0  //왼쪽

	dy := make([]int, 4)
	dy[0] = 0  //아래
	dy[1] = 1  //오른쪽
	dy[2] = 0  //위
	dy[3] = -1 //왼쪽

	for i := 0; i < n; i++ {
		snail[i] = make([]int, n)
	}

	snail[x][y] = n * n

	for idx < 4 {
		if snail[x][y] == 1 {
			break
		}
		nx := x + dx[idx]
		ny := y + dy[idx]

		if nx >= 0 && ny >= 0 && nx < n && ny < n && snail[nx][ny] == 0 {
			snail[nx][ny] = snail[x][y] - 1
			if snail[nx][ny] == location {
				locationX = nx
				locationY = ny
			}
			if snail[nx][ny] == 1 {
				break
			}
			x = nx
			y = ny
		} else {
			idx++
		}

		if idx == 4 {
			idx = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, "%d ", snail[i][j])
		}
		fmt.Fprintln(w)
	}
	fmt.Fprint(w, locationX+1, locationY+1)

	w.Flush()

}
