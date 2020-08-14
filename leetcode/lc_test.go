package main

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(*testing.T) {
	a := []int{2, 4, 5, 3, 4, 2, 1}
	fmt.Println(LargestRectangleArea(a))
}

func TestMedianOfTwoSortedArrays(*testing.T) {
	a := []int{1, 2, 3, 6, 7}
	b := []int{4, 5, 8, 9}
	fmt.Println(MedianOfTwoSortedArrays(a, b))
	a = []int{1, 4}
	b = []int{2, 3}
	fmt.Println(MedianOfTwoSortedArrays(a, b))
}