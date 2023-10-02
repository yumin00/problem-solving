package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
관찰 횟수 : N
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscanln(r, &N) // 관찰 횟수 입력 받기

	cows := make([]int, 10) // 초기화 값 0

	for i := 0; i < 10; i++ {
		cows[i] = -1 // 소의 위치 값을 모두 -1 로 설정
	}

	cnt := 0

	for i := 0; i < N; i++ {
		var idx, location int // 0, 1 - left, right
		fmt.Fscan(r, &idx, &location)

		if cows[idx-1] == -1 { // 관찰 처음이면 indx에 location 추가
			cows[idx-1] = location
		} else if cows[idx-1] != -1 && cows[idx-1] != location { // 관찰한 값이 다르면, cnt + 1 하고 해당 location 값을 입력 받은 location 값으로 바꾸기
			cnt++
			cows[idx-1] = location
		} else if cows[idx-1] == location { // 관찰한 소의 위치가 전에 있던 소의 위치와 같으면 넘어가기
			continue
		}
	}

	fmt.Fprint(w, cnt)
	w.Flush()
}