package main

import "fmt"

func main() {
	u := int(4)
	p := &u
	v := int(5)
	q := &v
	
	h := new(int)
	*h = 6
	q = p 
	p = h
	fmt.Println(p, q, h)
}