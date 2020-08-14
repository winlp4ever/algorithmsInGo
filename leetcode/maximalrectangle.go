package main


func LargestRectangleArea(hs []int) int {
    hs = append(hs, 0)
    n := len(hs)
    st := [][]int{}
    mn := 0
    for i, v := range hs {
        k := i
        for len(st) > 0 && (v < st[len(st)-1][0] || i == n-1) {
            l := len(st)
            if _m := st[l-1][0] * (i-st[l-1][1]); mn < _m {
                mn = _m
            }
            k = st[l-1][1]
            st = st[:l-1]
		}
		st = append(st, []int{v, k})
    }
    return mn
}

func MaximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    mx := 0
    n := len(matrix[0])
    hiss := make([]int, n)
    for _, r := range matrix {
        row := []int{}
        for j, b := range r {
            row = append(row, int(b-'0') * (hiss[j]+1))
        }
        hiss = row
        if rec := LargestRectangleArea(hiss); rec > mx {
            mx = rec
        }
    }
    return mx
}