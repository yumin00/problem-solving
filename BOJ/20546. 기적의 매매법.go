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

/*
   문제

   준현 jun
   - 살 수 있다면 무조건 많이 산다.
   - 매수 가능 주식 수 : 현금 / 주가
   - 남은 현금 : 현금 - (매수 가능 주식수 * 주가)
   - 보유 주식 수 : 매수 가능 주식 수 ++

   성민 sung
   - 전량 매수, 전량 매도
   - 매수 가능 주식 수 : 현금 / 주가

   - 보유하고 있는 주식이 3일째 상승한다면, 전량 매도
   - 3일 연속 가격이 하락하는 주식은 다음날 상승한다고 가정 -> 전략 매수

   내기
   - 기간 : 1월 1일 ~ 1월 14일
   - 자산 : 현금 + 1월 14일의 주가 * 주식 수
*/
