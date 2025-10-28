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

func dfs(step int){
	for i := 0; i < n; i++ {
		if step == 1 || !used[i] && isSquare(res[len(res)-1] + i){
			res, used[i] = append(res, i), true
			dfs()
		}
	}
}

func SquareSumsRow(n int) []int{
	used = make([]bool, n+1)
	res = make([]int, 0)
	


	return nil
}