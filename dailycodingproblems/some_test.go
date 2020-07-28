package main

import (
	"fmt"
	"testing"
)

func TestMinPattern(*testing.T) {
	p := "()))(()((()()()))"
	sol := MinPattern(p)
	fmt.Println(sol, len(p), len(sol))
}

func TestBuyAndSellStocks(t *testing.T) {
	a := []int{1, 5, 2, 3, 1, 7, 3, 2, 2, 4}
	fmt.Println(BuyAndSellStocks(a))
}