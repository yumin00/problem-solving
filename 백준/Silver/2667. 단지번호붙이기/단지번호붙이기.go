package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var apart [][]bool
var n int
var apartCnt int
var houseCnt int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &n)

	apart = make([][]bool, n)

	for i := 0; i < n; i++ {
		apart[i] = make([]bool, n)
		for j := 0; j < n; j++ {
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscanf(r, "%1d", &a)
			if a == 1 {
				apart[i][j] = true
			} else {
				apart[i][j] = false
			}
		}
		fmt.Fscanln(r)
	}

	var cnt []int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if apart[i][j] == true {
				apart[i][j] = false
				houseCnt = bfs(i, j)
				if houseCnt > 0 {
					apartCnt++
					cnt = append(cnt, houseCnt)
				}
			}
		}
	}
	fmt.Fprintln(w, apartCnt)

	sort.Sort(sort.IntSlice(cnt))
	for i := 0; i < len(cnt); i++ {
		fmt.Fprintln(w, cnt[i])
	}
	w.Flush()
}

type Point struct {
	x int
	y int
}

func bfs(i, j int) (houseCnt int) {
	houseCnt = 1
	var q []Point
	q = append(q, Point{
		x: i,
		y: j,
	})

	directions := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for len(q) > 0 {
		for _, v := range directions {
			moveX := q[0].x + v[0]
			moveY := q[0].y + v[1]
			if moveX >= 0 && moveX < n && moveY >= 0 && moveY < n && apart[moveX][moveY] == true {
				houseCnt++
				q = append(q, Point{
					x: moveX,
					y: moveY,
				})
				apart[moveX][moveY] = false
			}
		}
		q = q[1:]
	}
	return houseCnt
}
