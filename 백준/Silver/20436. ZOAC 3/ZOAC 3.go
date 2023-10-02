package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type key struct {
	x int
	y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var left map[rune]key
	var right map[rune]key

	left = make(map[rune]key)
	right = make(map[rune]key)
	left['q'] = key{
		x: 0,
		y: 0,
	}
	left['w'] = key{
		x: 0,
		y: 1,
	}
	left['e'] = key{
		x: 0,
		y: 2,
	}
	left['r'] = key{
		x: 0,
		y: 3,
	}
	left['t'] = key{
		x: 0,
		y: 4,
	}
	left['a'] = key{
		x: 1,
		y: 0,
	}
	left['s'] = key{
		x: 1,
		y: 1,
	}
	left['d'] = key{
		x: 1,
		y: 2,
	}
	left['f'] = key{
		x: 1,
		y: 3,
	}
	left['g'] = key{
		x: 1,
		y: 4,
	}
	left['z'] = key{
		x: 2,
		y: 0,
	}
	left['x'] = key{
		x: 2,
		y: 1,
	}
	left['c'] = key{
		x: 2,
		y: 2,
	}
	left['v'] = key{
		x: 2,
		y: 3,
	}
	right['y'] = key{
		x: 0,
		y: 1,
	}
	right['u'] = key{
		x: 0,
		y: 2,
	}
	right['i'] = key{
		x: 0,
		y: 3,
	}
	right['o'] = key{
		x: 0,
		y: 4,
	}
	right['p'] = key{
		x: 0,
		y: 5,
	}
	right['h'] = key{
		x: 1,
		y: 1,
	}
	right['j'] = key{
		x: 1,
		y: 2,
	}
	right['k'] = key{
		x: 1,
		y: 3,
	}
	right['l'] = key{
		x: 1,
		y: 4,
	}
	right['b'] = key{
		x: 2,
		y: 0,
	}
	right['n'] = key{
		x: 2,
		y: 1,
	}
	right['m'] = key{
		x: 2,
		y: 2,
	}

	var alphabet1 rune
	var alphabet2 rune
	var word string

	fmt.Fscanf(r, "%c %c", &alphabet1, &alphabet2)
	fmt.Fscan(r, &word)

	var currentLX, currentLY int = left[(alphabet1)].x, left[(alphabet1)].y
	var currentRX, currentRY int = right[(alphabet2)].x, right[(alphabet2)].y

	ans := 0

	for _, v := range []rune(word) {
		if isVowel(v) {
			posRX := right[v].x
			posRY := right[v].y
			x := math.Abs(float64(currentRX - posRX))
			y := math.Abs(float64(currentRY - posRY))
			ans += int(x + y)
			currentRX = posRX
			currentRY = posRY
		} else {
			posLY := left[v].y
			posLX := left[v].x
			x := math.Abs(float64(currentLX - posLX))
			y := math.Abs(float64(currentLY - posLY))
			ans += int(x + y)
			currentLX = posLX
			currentLY = posLY
		}
		ans++
	}
	fmt.Fprint(w, ans)
	w.Flush()
}

func isVowel(alphabet rune) bool {
	switch alphabet {
	case 'y', 'u', 'i', 'o', 'p', 'h', 'j', 'k', 'l', 'b', 'n', 'm':
		return true
	default:
		return false
	}
}
