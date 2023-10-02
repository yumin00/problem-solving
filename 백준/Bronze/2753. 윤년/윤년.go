package main

import (
	"fmt"
)

func mod(int1, int2 int) (int, error) {

	if int2 > 0 {
		return int1 % int2, nil
	} else {
		return 0, fmt.Errorf("divsion by zero")
	}
}

func main() {
	var year int

	fmt.Scan(&year)

	leapYear1, _ := mod(year, 4)
	leapYear2, _ := mod(year, 100)
	leapYear3, _ := mod(year, 400)

	if leapYear1 == 0 && leapYear2 > 0 {
		fmt.Println(1)
	} else if leapYear3 == 0 {
		fmt.Println(1)
	} else {
	fmt.Println(0)
	}
}