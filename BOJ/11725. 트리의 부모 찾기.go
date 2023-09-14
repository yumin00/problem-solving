package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []interface{}

func (q *Queue) EnQueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	var tree [][]int
	tree = make([][]int, n+1)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscanln(r, &x, &y)
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}

	var parents []int
	parents = make([]int, n+1)
	parents[1] = 1

	q := Queue{}
	q.EnQueue(1)

	for q.IsEmpty() == false {
		start := q.DeQueue().(int)
		for _, v := range tree[start] {
			node := v
			if parents[node] == 0 {
				parents[node] = start
				q.EnQueue(node)
			}
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintln(w, parents[i])
	}
	w.Flush()
}

//type Queue []interface{}
//
//// IsEmpty - 큐가 비어있는지 확인하는 함수.
//func (q *Queue) IsEmpty() bool {
//	return len(*q) == 0
//}
//
//// Enqueue - 큐에 값을 추가하는 함수.
//func (q *Queue) Enqueue(data interface{}) {
//	*q = append(*q, data) // 큐 끝에 값을 추가함.
//}
//
//// Dequeue - 큐에 첫번째 요소를 반환하고 제거하는 함수.
//func (q *Queue) Dequeue() interface{} {
//	if q.IsEmpty() {
//		fmt.Println("queue is empty")
//		return nil
//	}
//	data := (*q)[0] // 큐에 첫번째 값을 가져옴.
//	*q = (*q)[1:]   // 큐에 첫번째 데이터를 제거함.
//	return data
//}
//
//func main() {
//	r := bufio.NewReader(os.Stdin)
//	w := bufio.NewWriter(os.Stdout)
//	var n int
//	fmt.Fscanln(r, &n)
//
//	var neighbor [][]int
//	var parents []int
//	neighbor = make([][]int, n+1)
//	parents = make([]int, n+1)
//	parents[1] = 1
//
//	for i := 0; i < n-1; i++ {
//		var a, b int
//		fmt.Fscanln(r, &a, &b)
//		neighbor[a] = append(neighbor[a], b)
//		neighbor[b] = append(neighbor[b], a)
//	}
//
//	problemQ := Queue{}
//	problemQ.Enqueue(1)
//
//	for !problemQ.IsEmpty() {
//		start := problemQ.Dequeue().(int)
//		for _, node := range neighbor[start] {
//			if parents[node] == 0 {
//				parents[node] = start
//				problemQ.Enqueue(node)
//			}
//		}
//	}
//
//	for i := 2; i <= n; i++ {
//		fmt.Fprintln(w, parents[i])
//	}
//	w.Flush()
//}

//func main() {
//	r := bufio.NewReader(os.Stdin)
//	w := bufio.NewWriter(os.Stdout)
//
//	var N int
//	fmt.Fscan(r, &N)
//
//	node := make([][2]int, N-1)
//	visited := make([]bool, N-1)
//	var q [][]int
//
//	for i := 0; i < N-1; i++ {
//		var x, y int
//		fmt.Fscan(r, &x, &y)
//		node[i][0] = x
//		node[i][1] = y
//		if x == 1 {
//			visited[i] = true
//			q = append(q, []int{1, y})
//		} else if y == 1 {
//			visited[i] = true
//			q = append(q, []int{1, node[i][0]})
//		} else {
//			continue
//		}
//	}
//
//	lenghtQ := len(q)
//
//	for i := 0; ; i++ {
//		if i >= lenghtQ {
//			break
//		}
//		for j := 0; j < N-1; j++ {
//			if node[j][0] == q[i][1] && visited[j] == false {
//				visited[j] = true
//				q = append(q, []int{node[j][0], node[j][1]})
//			}
//			if node[j][1] == q[i][1] && visited[j] == false {
//				visited[j] = true
//				q = append(q, []int{node[j][1], node[j][0]})
//			}
//		}
//		lenghtQ = len(q)
//	}
//
//	var result []int
//	for i := 2; i < N+1; i++ {
//		for j := 0; j < lenghtQ; j++ {
//			if q[j][1] == i {
//				result = append(result, q[j][0])
//			}
//		}
//	}
//	for i := 0; i < len(result); i++ {
//		fmt.Fprintln(w, result[i])
//	}
//	w.Flush()
//}

//func main() {
//	r := bufio.NewReader(os.Stdin)
//	w := bufio.NewWriter(os.Stdout)
//
//	var N int
//	fmt.Fscan(r, &N)
//
//	node := make([][2]int, N-1)
//	visited := make([]bool, N-1)
//
//	for i := 0; i < N-1; i++ {
//		var x, y int
//		fmt.Fscan(r, &x, &y)
//		node[i][0] = x
//		node[i][1] = y
//	}
//
//	var q [][]int
//
//	for i := 0; i < N-1; i++ {
//		if node[i][0] == 1 {
//			visited[i] = true
//			q = append(q, []int{1, node[i][1]})
//		}
//		if node[i][1] == 1 {
//			visited[i] = true
//			q = append(q, []int{1, node[i][0]})
//		}
//	}
//
//	lenghtQ := len(q)
//
//	for i := 0; ; i++ {
//		if i >= lenghtQ {
//			break
//		}
//		for j := 0; j < N-1; j++ {
//			if node[j][0] == q[i][1] && visited[j] == false {
//				visited[j] = true
//				q = append(q, []int{node[j][0], node[j][1]})
//			}
//			if node[j][1] == q[i][1] && visited[j] == false {
//				visited[j] = true
//				q = append(q, []int{node[j][1], node[j][0]})
//			}
//		}
//		lenghtQ = len(q)
//	}
//
//	var result []int
//	for i := 2; i < N+1; i++ {
//		for j := 0; j < lenghtQ; j++ {
//			if q[j][1] == i {
//				result = append(result, q[j][0])
//			}
//		}
//	}
//	fmt.Fprintln(w, result)
//	w.Flush()
//}

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
