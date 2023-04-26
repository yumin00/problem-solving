package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
8진수가 주어졌을 때, 2진수로 변환하는 프로그램을 작성하시오.

첫째 줄에 8진수가 주어진다. 주어지는 수의 길이는 333,334을 넘지 않는다.

첫째 줄에 주어진 수를 2진수로 변환하여 출력한다. 수가 0인 경우를 제외하고는 반드시 1로 시작해야 한다.

*/

/*
	1. 8진수 입력
	ex) 256
	2. 각 숫자 하나씩 떼기
	3. for문으로 숫자를 2로 나눠서 몫과 나머지
*/

func main() {
	var octal string
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscan(r, &octal) // 1. 8진수 입력

	for i, o := range octal {
		octalByte := byte(o)

		demical, _ := strconv.ParseInt(string(octalByte), 8, 8)
		binary := strconv.FormatInt(demical, 2)

		if i == 0 {
			fmt.Fprint(w, binary)
		} else {
			fmt.Fprintf(w, "%03s", binary)
		}
	}
	fmt.Fprintln(w)
	w.Flush()
}

//for i := 0; ; i++ {
//	octalByte, err := r.ReadByte()
//	if err != nil {
//		break
//	}
//	//octalByte := byte(octal[i])
//	demical, _ := strconv.ParseInt(string(octalByte), 8, 8)
//	binary := strconv.FormatInt(demical, 2)
//
//	if i == 0 {
//		fmt.Print(binary)
//	} else {
//		fmt.Printf("%03s", binary)
//	}
//}
