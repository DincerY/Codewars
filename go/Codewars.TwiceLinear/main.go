package main

import (
	"sort"
)

func main() {
	DblLinear(10)
	DblLinear(50)
	DblLinear(100)

}

// It worked but we got timed out error
func DblLinear(n int) int {
	res := []int{1}
	dic := make(map[int]struct{})
	dic[1] = struct{}{}

	for true {
		i := res[0]
		res = res[1:]
		res = append(res, (i*2)+1)
		res = append(res, (i*3)+1)
		sort.Ints(res)

		if _, exists := dic[i]; exists {
			continue
		}
		if len(dic) == n+1 {
			break
		}
		dic[i] = struct{}{}
	}
	max := 0

	for k := range dic {
		if k > max {
			max = k
		}
	}
	return max
}
