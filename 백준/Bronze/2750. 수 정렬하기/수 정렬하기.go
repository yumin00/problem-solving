package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
N개의 수가 주어졌을 때, 이를 오름차순으로 정렬하는 프로그램을 작성하시오.

입력
첫째 줄에 수의 개수 N(1 ≤ N ≤ 1,000)이 주어진다. 둘째 줄부터 N개의 줄에는 수가 주어진다. 이 수는 절댓값이 1,000보다 작거나 같은 정수이다. 수는 중복되지 않는다.

출력
첫째 줄부터 N개의 줄에 오름차순으로 정렬한 결과를 한 줄에 하나씩 출력한다.
*/

/*
인접한 것들 끼리 비교해서 큰 값을 오른쪽으로 작은 값을 왼쪽으로
한 번을 반복을 하면 맨 오른쪽에 제일 큰 값이 남음
제일 큰 값을 제외하고 다시 반복
시간 : n(n-1)/2 => O(n^2)
*/

func main() {
	//수의 개수 받기
	var count int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &count)

	arr := make([]int, count)

	//수 받기
	for i := 0; i < count; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := 0; i < count-1; i++ {
		for j := 0; j < count-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < count; i++ {
		fmt.Println(arr[i])
	}
}
