package main

import (
	"fmt"
)

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