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

func Test208(t *testing.T) {
	a := []int{1, 3, 2, 6, 5, 4, 7}
	r := Slice2List(a)
	fmt.Println(r.ToString())
	b := Partition(r, 5)
	fmt.Println(b.ToString())
}

func TestCheckBipartie(t *testing.T) {
	//edges := [][]int{{1, 5}, {2, 4}, {2, 5}, {1, 6}, {3, 5}}
	edges := [][]int{{1, 2}, {3, 2}}
	g := NewGraph(edges)
	fmt.Println(g)
	fmt.Println(g.IsBipartie())
}

type T struct {
	Val int
}

func TestNil(t *testing.T) {
	var a *T = nil
	a = new(T)
	fmt.Println(a)
}