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
	var N int
	fmt.Fscanln(r, &N)

	game := make([]string, N)
	play := make([]string, N)
	trap := make([][]bool, N)
	visited := make([][]bool, N)
	cnt := 0

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &game[i])
		trap[i] = make([]bool, N)
		r := []rune(game[i])
		for idx, val := range r {
			if string(val) == "*" {
				trap[i][idx] = true
			} else {
				trap[i][idx] = false
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &play[i])
	}

	answer := make([]string, N)
	lose := false
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
		str := ""
		r := []rune(play[i])
		for idx, val := range r {
			if string(val) == "x" {
				visited[i][idx] = true
				cnt = 0
				if trap[i][idx] == true {
					lose = true
					str += "."
				} else {
					if i > 0 && idx > 0 && trap[i-1][idx-1] == true {
						cnt++
					}
					if i > 0 && trap[i-1][idx] == true {
						cnt++
					}
					if i > 0 && idx+1 < N && trap[i-1][idx+1] == true {
						cnt++
					}
					if idx > 0 && trap[i][idx-1] == true {
						cnt++
					}
					if idx+1 < N && trap[i][idx+1] == true {
						cnt++
					}
					if idx > 0 && i+1 < N && trap[i+1][idx-1] == true {
						cnt++
					}
					if i+1 < N && trap[i+1][idx] == true {
						cnt++
					}
					if i+1 < N && idx+1 < N && trap[i+1][idx+1] == true {
						cnt++
					}

					str += strconv.Itoa(cnt)
				}
			} else {
				visited[i][idx] = false
				str += "."
			}
		}
		answer[i] = str
	}
	if lose {
		for i := 0; i < N; i++ {
			r := []rune(answer[i])
			for idx, val := range r {
				if trap[i][idx] {
					fmt.Fprint(w, "*")
					w.Flush()
				} else {
					fmt.Fprint(w, string(val))
					w.Flush()
				}
			}
			fmt.Fprintln(w)
			w.Flush()
		}
	} else {
		for j := 0; j < N; j++ {
			fmt.Fprintln(w, answer[j])
			w.Flush()
		}
	}
}
