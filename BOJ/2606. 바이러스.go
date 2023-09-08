package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []interface{}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var count int
	var n int
	fmt.Fscanln(r, &count)
	fmt.Fscanln(r, &n)

	computer := make([][2]int32, n)
	check := make([]bool, n)

	for i := 0; i < n; i++ {
		var a, b int32
		fmt.Fscanln(r, &a, &b)
		computer[i][0] = a
		computer[i][1] = b
	}

	for i := 0; i < n; i++ {
		check[i] = false
	}

	var q Queue

	for i := 0; i < n; i++ {
		if computer[i][0] == 1 {
			q = append(q, computer[i][1])
			check[i] = true
		} else if computer[i][1] == 1 {
			q = append(q, computer[i][0])
			check[i] = true
		} else {
			check[i] = false
			continue
		}
	}

	for i := 0; ; i++ {
		lengthQ := len(q)
		if i >= lengthQ || lengthQ == 0 {
			break
		}

		for j := 0; j < n; j++ {
			if check[j] == true {
				continue
			}
			if computer[j][0] == q[i] {
				isDuplicate := false
				for v := 0; v < lengthQ; v++ {
					if computer[j][1] == q[v] {
						isDuplicate = true
						break
					}
				}
				if isDuplicate == false {
					q = append(q, computer[j][1])
					check[j] = true
				}
			} else if computer[j][1] == q[i] {
				isDuplicate := false
				for v := 0; v < lengthQ; v++ {
					if computer[j][0] == q[v] {
						isDuplicate = true
						break
					}
				}
				if isDuplicate == false {
					q = append(q, computer[j][0])
					check[j] = true
				}
			}
		}
	}
	fmt.Fprintln(w, len(q))
	w.Flush()
}
