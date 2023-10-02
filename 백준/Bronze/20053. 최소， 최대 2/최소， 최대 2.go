package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//func main() {
//	r := bufio.NewReader(os.Stdin)
//	var testCase int
//
//	fmt.Fscan(r, &testCase)
//
//	//results := make([]string, testCase) // 모든 결과를 저장할 슬라이스 만들기
//
//	for t := 0; t < testCase; t++ {
//		var countN int
//		var max int
//		var min int
//		fmt.Fscan(r, &countN) //2. 몇 개의 숫자를 받을 것인가?
//
//		var intN int
//		for i := 0; i < countN; i++ {
//			fmt.Fscan(r, &intN) //3. 숫자 입력 받기
//			if i == 0 {
//				max = intN
//				min = intN
//			} else {
//				if intN > max {
//					max = intN
//				}
//				if intN < min {
//					min = intN
//				}
//			}
//		}
//		fmt.Printf("%d %d\n", min, max)
//	}
//}

// 시간 오류....
func main() {
	r := bufio.NewReader(os.Stdin)
	var testCase int

	fmt.Fscan(r, &testCase) // 1. 테스트를 몇 번 할 것인가?

	results := make([]string, testCase) // 모든 결과를 저장할 슬라이스 만들기

	for i := 0; i < testCase; i++ {
		var countN int
		fmt.Fscan(r, &countN) //2. 몇 개의 숫자를 받을 것인가?

		var intN = make([]int, countN)
		for i := 0; i < countN; i++ {
			fmt.Fscan(r, &intN[i]) //3. 숫자 입력 받기
		}

		sort.Slice(intN, func(i, j int) bool { //4. slice 정렬
			return intN[i] < intN[j]
		})
		// 최대값, 최소값 문자열로 저장
		results[i] = fmt.Sprintf("%d %d", intN[0], intN[len(intN)-1])
	}

	// 모든 결과 출력
	for _, result := range results {
		fmt.Println(result)
	}
}