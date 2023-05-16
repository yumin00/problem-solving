package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N, M int

	fmt.Fscanln(r, &N, &M) // 1. 전구 개수, 명령어 개수 입력

	status := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &status[i]) // 2. N개의 전구 상태
	}

	command := make([][3]int, M)

	for i := 0; i < M; i++ {
		for j := 0; j < 3; j++ {
			fmt.Fscan(r, &command[i][j]) //3. 명령어 입력
		}
		idx := command[i][1]
		x := command[i][2]

		if command[i][0] == 1 {
			status[idx-1] = x
		} else if command[i][0] == 2 {
			for i := idx - 1; i < x; i++ {
				if status[i] == 0 {
					status[i] = 1
				} else {
					status[i] = 0
				}
			}
		} else if command[i][0] == 3 {
			for i := idx - 1; i < x; i++ {
				if status[i] == 1 {
					status[i] = 0
				}
			}
		} else if command[i][0] == 4 {
			for i := idx - 1; i < x; i++ {
				if status[i] == 0 {
					status[i] = 1
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		fmt.Fprint(w, status[i])
		if i != N-1 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w)
		w.Flush()
	}
}

/*
   입력

   1 : N(전구 개수) M(명령어 개수)
   2 : N개의 전구 상태 ex) 010101
   3
   a - 몇번째 명령어
   b, c - a가 1이면 각각 i , x
   b,c  - a가 2,3,4 중 하나면 l, r

   출력
   전구가 어떤 상태인지 출력
*/
/*
   문제

   N개의 전구가 있고 맨 왼쪽에 있는 전구를 첫 번째라고 하자. 전구의 상태는 두 가지가 있으며 이를 숫자로 표현한다.

   1은 전구가 켜져 있는 상태를 의미하고,
   0은 전구가 꺼져 있는 상태를 의미한다.

   전구를 제어하는 명령어가 1번부터 4번까지 4개가 있다. 아래 표는 각 명령어에 대한 설명이다.

   1번 : [i x] : i번째 전구를 x 상태로
   2번 : [l, r] : l~r 0 -> 1, 1 -> 0
   3번 : [l, r] : l~r 끄기
   4번 : [l, r] : l~r 켜기

*/
