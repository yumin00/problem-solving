package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(r, &N)

	node := make([][2]int, N-1)
	visited := make([]bool, N-1)

	for i := 0; i < N-1; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		node[i][0] = x
		node[i][1] = y
	}

	var q [][]int

	for i := 0; i < N-1; i++ {
		if node[i][0] == 1 {
			visited[i] = true
			q = append(q, []int{1, node[i][1]})
		}
		if node[i][1] == 1 {
			visited[i] = true
			q = append(q, []int{1, node[i][0]})
		}
	}

	lenghtQ := len(q)

	for i := 0; ; i++ {
		if i >= lenghtQ {
			break
		}
		for j := 0; j < N-1; j++ {
			if node[j][0] == q[i][1] && visited[j] == false {
				visited[j] = true
				q = append(q, []int{node[j][0], node[j][1]})
			}
			if node[j][1] == q[i][1] && visited[j] == false {
				visited[j] = true
				q = append(q, []int{node[j][1], node[j][0]})
			}
		}
		lenghtQ = len(q)
	}

	var result []int
	for i := 2; i < N+1; i++ {
		for j := 0; j < lenghtQ; j++ {
			if q[j][1] == i {
				result = append(result, q[j][0])
			}
		}
	}
	fmt.Fprintln(w, result)
	w.Flush()
}

//func main() {
//	r := bufio.NewReader(os.Stdin)
//	w := bufio.NewWriter(os.Stdout)
//
//	var N int
//	fmt.Fscan(r, &N)
//
//	node := make([][]int, N+1)
//
//	for i := range node {
//		node[i] = make([]int, N+1)
//	}
//
//	for i := 0; i < N-1; i++ {
//		var x, y int
//		fmt.Fscan(r, &x, &y)
//		node[x][y] = 1
//		node[y][x] = 1
//	}
//
//	for i := 2; i < N+1; i++ {
//		for j := 0; j < len(node[i]); j++ {
//			if node[i][j] == 1 {
//				fmt.Fprintln(w, j)
//				break
//			} else {
//				continue
//			}
//		}
//	}
//	w.Flush()
//}
