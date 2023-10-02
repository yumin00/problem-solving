package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []interface{}

// IsEmpty - 큐가 비어있는지 확인하는 함수.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Enqueue - 큐에 값을 추가하는 함수.
func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data) // 큐 끝에 값을 추가함.
}

// Dequeue - 큐에 첫번째 요소를 반환하고 제거하는 함수.
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	data := (*q)[0] // 큐에 첫번째 값을 가져옴.
	*q = (*q)[1:]   // 큐에 첫번째 데이터를 제거함.
	return data
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(r, &n)
	var neighbor [][]int
	var parents []int
	neighbor = make([][]int, n+1)
	parents = make([]int, n+1)
	parents[1] = 1
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		neighbor[a] = append(neighbor[a], b)
		neighbor[b] = append(neighbor[b], a)
	}

	problemQ := Queue{}
	problemQ.Enqueue(1)

	for !problemQ.IsEmpty() {
		start := problemQ.Dequeue().(int)
		for _, v := range neighbor[start] {
			node := v
			if parents[node] == 0 {
				parents[node] = start
				problemQ.Enqueue(node)
			}
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintln(w, parents[i])
	}
	w.Flush()
}