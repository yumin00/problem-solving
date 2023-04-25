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

/*
N개의 정수가 주어진다. 이때, 최솟값과 최댓값을 구하는 프로그램을 작성하시오.

[입력]
첫째 줄에 테스트 케이스의 개수 T (1 ≤ T ≤ 10)가 주어진다. 각 테스트 케이스는 두 줄로 이루어져 있다.

각 테스트 케이스의 첫째 줄에 정수의 개수 N (1 ≤ N ≤ 1,000,000)이 주어진다. 둘째 줄에는 N개의 정수를 공백으로 구분해서 주어진다. 모든 정수는 -1,000,000보다 크거나 같고, 1,000,000보다 작거나 같은 정수이다.

[출력]
각 테스트 케이스마다 주어진 정수 N개의 최솟값과 최댓값을 공백으로 구분해 한 줄에 하나씩 차례대로 출력한다.
*/

/*
1. 첫째줄 : 테스트 케이스의 개수 T scan
2. 테스트 케이스 개수 T만큼 countN을 받아야함

1. countN 만큼의 intN을 받아야함
if countN만큼의 intN length를 안 받으면 다시 입력하라는 메세지 만들기
2. 입력 : countN 개수 만큼의 숫자 입력
3. 결과값 저장
3. 출력 : 숫자의 max, min 출력

*/
