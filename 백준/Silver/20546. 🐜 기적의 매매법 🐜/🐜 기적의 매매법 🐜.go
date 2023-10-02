package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var seedMoney int32
	fmt.Fscanln(r, &seedMoney)

	dailyStock := make([]int32, 14)
	for i := 0; i < 14; i++ {
		fmt.Fscan(r, &dailyStock[i])
	}

	//준현
	var junStock, junStockCount int32
	junMoney := seedMoney
	for i := 0; i < 14; i++ {
		if junMoney >= dailyStock[i] {
			junStock = junMoney / dailyStock[i]
			junMoney = junMoney - junStock*dailyStock[i]
			junStockCount = junStockCount + junStock
		} else {
			continue
		}
	}
	junFinalMoney := junMoney + dailyStock[13]*junStockCount

	var sungStock, sungStockCount int32
	sungMoney := seedMoney

	for i := 3; i < 14; i++ {
		if dailyStock[i] < dailyStock[i-1] && dailyStock[i-1] < dailyStock[i-2] && dailyStock[i-2] < dailyStock[i-3] && sungMoney > dailyStock[i] {
			sungStock = sungMoney / dailyStock[i]
			sungMoney = sungMoney - sungStock*dailyStock[i]
			sungStockCount = sungStockCount + sungStock
		} else if dailyStock[i] > dailyStock[i-1] && dailyStock[i-1] > dailyStock[i-2] && dailyStock[i-2] > dailyStock[i-3] && sungStockCount != 0 {
			sungMoney = sungMoney + sungStockCount*dailyStock[i]
			sungStockCount = 0
		} else {
			continue
		}
	}
	sungFinalMoney := sungMoney + dailyStock[13]*sungStockCount

	if junFinalMoney > sungFinalMoney {
		fmt.Fprint(w, "BNP")
		w.Flush()
	} else if junFinalMoney < sungFinalMoney {
		fmt.Fprint(w, "TIMING")
		w.Flush()
	} else {
		fmt.Fprint(w, "SAMESAME")
		w.Flush()
	}

}