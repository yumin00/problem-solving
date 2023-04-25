package main

import (
	"fmt"
)

const INVALID = -1

func mod2(int1, int2 int) int {

	if int2 > 0 {
		return int1 % int2
	} else {
		return INVALID
	}
}

func main() {
	var year int

	fmt.Scan(&year)

	leapYear1 := mod2(year, 4)
	leapYear2 := mod2(year, 100)
	leapYear3 := mod2(year, 400)

	if leapYear1 == 0 && leapYear2 > 0 || leapYear3 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

/*
연도가 주어졌을 때, 윤년이면 1, 아니면 0을 출력하는 프로그램을 작성하시오.

윤년은 연도가 4의 배수이면서, 100의 배수가 아닐 때 또는 400의 배수일 때이다.

예를 들어, 2012년은 4의 배수이면서 100의 배수가 아니라서 윤년이다. 1900년은 100의 배수이고 400의 배수는 아니기 때문에 윤년이 아니다. 하지만, 2000년은 400의 배수이기 때문에 윤년이다.
*/

/*
윤년 ==
year을 4로 나눴을 때 나머지가 0인 경우 && year을 100으로 나눴을 때 나머지가 0 초과인 경우
or
year을 400으로 나눴을 때 나머지가 0인 경우
*/
