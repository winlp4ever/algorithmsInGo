package main

import "strconv"

type TrieNode struct {
	Letter rune
	End bool
	Childs [](*TrieNode)
}

type Triple struct {
	H int
	N *TrieNode
	Sym string
}

func (root *TrieNode) ToString() string {
	s := ""
	Q := []Triple{{0, root, "#"}}
	h := 0
	for len(Q) > 0 {
		p := Q[0]
		Q = Q[1:]
		if p.H > h {
			h ++
			s += "\n"
		}
		s += p.Sym + "-" + strconv.Itoa(p.H) + " "
		for _, n := range p.N.Childs {
			if n != nil {
				Q = append(Q, Triple{p.H + 1, n, string(p.N.Letter) + "->" + string(n.Letter)})
			}
		}
	}
	return s
}

func addString(root *TrieNode, s string) *TrieNode {
	r := root
	for i, c := range s {
		if r.Childs[int(c - 'a')] == nil {
			r.Childs[int(c - 'a')] = new(TrieNode)
			r.Childs[int(c - 'a')].Letter = c
			r.Childs[int(c - 'a')].Childs = make([](*TrieNode), 26)
			if i == len(s)-1 {
				r.Childs[int(c - 'a')].End = true
			}
		}
		r = r.Childs[int(c - 'a')]
	}
	return root
}

func GenTrie(words []string) *TrieNode {
	root := new(TrieNode)
	root.End = false
	root.Childs = make([](*TrieNode), 26)
	root.Letter = '#'
	for _, w := range words {
		root = addString(root, w)
	}
	return root
}

