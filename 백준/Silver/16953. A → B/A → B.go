package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b int
	fmt.Fscanf(r, "%d %d", &a, &b)

	var cnt int

	for {
		if a == b {
			cnt++
			break
		} else if a > b {
			cnt = -1
			break
		}

		if b%2 == 0 {
			b = b / 2
			cnt++
		} else if getLastDigit(b) == 1 {
			b = (b - 1) / 10
			cnt++
		} else {
			cnt = -1
			break
		}
	}
	fmt.Fprintln(w, cnt)
	w.Flush()
}

func getLastDigit(number int) int {
	numStr := strconv.Itoa(number)

	lastDigit, err := strconv.Atoi(string(numStr[len(numStr)-1]))
	if err != nil {
		return 0
	}
	return lastDigit
}
