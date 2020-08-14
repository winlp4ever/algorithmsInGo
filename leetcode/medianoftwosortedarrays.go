package main

import "fmt"

func MedianOfTwoSortedArrays(a []int, b []int) (int, int) {
	fmt.Println(a, b)
	m, n := len(a), len(b)
	ai := a[m / 2]
	bj := b[n / 2]
	if m == 1 && n == 1 {
		return ai, bj
	}
	if ai > bj {
		return MedianOfTwoSortedArrays(a[:(m+1)/2], b[n/2:])
	} 
	return MedianOfTwoSortedArrays(a[m/2:], b[:(n+1)/2])
}