package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i {
				fmt.Fprint(w, "*")
			} else {
				fmt.Fprint(w, "")
			}
		}
		fmt.Fprintln(w)
		w.Flush()
	}
}