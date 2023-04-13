package main

import (
	"fmt"
)

/*
교실엔 학생이 30명이 있는데, 학생 명부엔 각 학생별로 1번부터 30번까지 출석번호가 붙어 있다.
교수님이 내준 특별과제를 28명이 제출했는데, 그 중에서 제출 안 한 학생 2명의 출석번호를 구하는 프로그램을 작성하시오.
*/

/*

학생 : 1번 ~ 30번
28명이 제출하고 2명은 제출 안 함 -> 28개만 받으면 됨 -< i 는 28까지


for문으로 28개의 숫자를 받기
받은 숫자를 순서대로


*/

func main() {
	studentNumber := make([]bool, 30)

	for i := 0; i < 28; i++ {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Errorf("Scan error")
		}
		studentNumber[n-1] = true
	}

	for i := 0; i < 30; i++ {
		if studentNumber[i] == false {
			fmt.Println(i + 1)
		}
	}

}
