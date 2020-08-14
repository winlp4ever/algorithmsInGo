package main

import (
	"strconv"
)

type Node struct {
	Val int
	Next *Node
}

func Slice2List(nums []int) *Node {
	n := len(nums)
	root := new(Node)
	node := root
	for i, u := range nums {
		node.Val = u
		if i < n-1 {
			node.Next = new(Node)
			node = node.Next
		}
	}
	return root
}

func (n *Node) ToString() string {
	s := ""
	node := n
	for node != nil {
		s += strconv.Itoa(node.Val)
		if node.Next != nil {
			s += " -> "
		}
		node = node.Next
	}
	return s
}

func Partition(root *Node, k int) *Node {
	u := root
	p := root
	for p != nil {
		if (p.Val < k) {
			p.Val, u.Val = u.Val, p.Val
			if u.Next != nil {
				u = u.Next
			} else {
				break
			}
		}
		p = p.Next
	}
	return root
}