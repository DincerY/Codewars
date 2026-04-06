package main

import (
	"sort"
)

func main() {
	DblLinear(1000)

	DblLinear(10)
	DblLinear(50)
	DblLinear(100)

}

// It worked but we got timed out error
func DblLinear(n int) int {
	res := []int{1}
	dic := make(map[int]struct{})
	max := 0

	for len(dic) != n+1 {
		i := res[0]
		res = res[1:]
		if _, exists := dic[i]; exists {
			continue
		}
		res = append(res, (i*2)+1)
		res = append(res, (i*3)+1)
		sort.Ints(res)

		dic[i] = struct{}{}
		if i > max {
			max = i
		}
	}

	return max
}
