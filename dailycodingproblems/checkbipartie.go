package main

type Graph struct {
	Adjs map[int][]int
}

func NewGraph(edges [][]int) *Graph {
	g := new(Graph)
	g.Adjs = map[int][]int{}
	for _, e := range edges {
		if _, ok := g.Adjs[e[0]]; ok {
			g.Adjs[e[0]] = append(g.Adjs[e[0]], e[1])
		} else {
			g.Adjs[e[0]] = []int{e[1]}
		}
		if _, ok := g.Adjs[e[1]]; ok {
			g.Adjs[e[1]] = append(g.Adjs[e[1]], e[0])
		} else {
			g.Adjs[e[1]] = []int{e[0]}
		}
	}
	return g
}

func (g *Graph) IsBipartie() bool {
	partitions := map[int]int{}
	for k, _ := range g.Adjs {
		_, ok := partitions[k]
		if ok {
			continue
		} 
		Q := [][]int{{0, k}}
		for len(Q) > 0 {
			v := Q[len(Q)-1]
			Q = Q[:len(Q)-1]
			for _, u := range g.Adjs[v[1]] {
				if p, ok := partitions[u]; ok {
					if p != 1-v[0] {
						return false
					} 
				} else {
					Q = append(Q, []int{1-v[0], u})
					partitions[u] = 1-v[0]
				}
			}
		}
	}
	return true
}