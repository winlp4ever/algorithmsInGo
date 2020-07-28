package main

import (
	"fmt"
	"strings"
)

func MinPattern(p string) string {
	mn := 0
	d := 0
	for _, c := range p {
		if c == '(' {
			d++
		} else {
			d--
		}
		if mn > d {
			mn = d
		}
	}
	//fmt.Println(mn, d)
	return strings.Repeat("(", -mn) + p + strings.Repeat(")", -mn + d)
}

func TestMinPattern() {
	p := "()))(()((()()()))"
	sol := MinPattern(p)
	fmt.Println(sol, len(p), len(sol))
}

func main() {
	TestMinPattern()
}



