package main

import "math"

func main() {

	SquareSumsRow(10)
}

var (
	found bool
	res   []int
	used  []bool
)

func isSquare(val int) bool {
	sqrt := int(math.Sqrt(float64(val)))
	if val > 0 && sqrt*sqrt == val {
		return true
	}
	return false
}

func dfs(step, n int){
	if step == n+1 {
		found = true
		return
	}
	for i := 1; i <= n; i++ {
		if step == 1 || !used[i] && isSquare(res[len(res)-1] + i){
			res, used[i] = append(res, i), true
			dfs(step+1,n)
			if found {
				return
			}
			res, used[i] = res[:len(res)-1], false
		}
	}
}

func SquareSumsRow(n int) []int{
	found = false
	used = make([]bool, n+1)
	res = make([]int, 0)
	dfs(1, n)
	if len(res) != n {
		return nil
	}
	return res
}