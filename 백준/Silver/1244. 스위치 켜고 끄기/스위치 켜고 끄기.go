package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N int
	var studentCount int

	fmt.Fscan(r, &N)

	s := make([]bool, N)

	for i := 0; i < N; i++ {
		var temp int
		fmt.Fscan(r, &temp)
		if temp == 0 {
			s[i] = true
		} else {
			s[i] = false
		}
	}

	fmt.Fscan(r, &studentCount)

	for i := 0; i < studentCount; i++ {
		var sex, location int
		fmt.Fscan(r, &sex, &location)
		if sex == 1 {
			for j := 1; j <= N; j++ {
				if location*j-1 >= N {
					break
				} else {
					s[location*j-1] = !s[location*j-1]
				}
			}
		} else if sex == 2 {
			cnt := 0
			for j := 1; j < N; j++ {
				if location-1-j >= 0 && location+j-1 <= N-1 && s[location-1-j] == s[location-1+j] {
					cnt++
				} else {
					for idx := location - 1 - cnt; idx < location+cnt; idx++ {
						s[idx] = !s[idx]
					}
					break
				}
			}
		}
	}

	result := make([]int, N)

	for i := 0; i < N; i++ {
		if s[i] == true {
			result[i] = 0
		} else {
			result[i] = 1
		}
	}

	//for i := 0; i < N; i++ {
	//	if i%20 == 19 {
	//		fmt.Fprintln(w, result[i])
	//	} else {
	//		fmt.Fprintf(w, "%d ", result[i])
	//	}
	//}

	for i := 0; i < N; i++ {
		if i != 0 && i%20 == 0 {
			fmt.Fprintln(w)
		}

		fmt.Fprintf(w, "%d ", result[i])
	}

	w.Flush()
}
