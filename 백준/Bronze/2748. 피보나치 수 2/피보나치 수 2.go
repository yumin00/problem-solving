package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var result int

	if n == 0 {
		result = 0
	} else if n <= 2 {
		result = 1
	}

	var index int
	last1 := 1
	last2 := 1

	for index = 2; index < n; index++ {
		result = last1 + last2
		last1 = last2
		last2 = result
	}

	fmt.Println(result)
}