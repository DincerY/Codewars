package main

import "math"

func main() {
	SquareSumsRow(10)
}

func SquareSumsRow(n int) []int {
	res := make([]int, 0, n)
	usedMap := make(map[int]struct{})
	if DFS(&res,usedMap,n) {
		return res
	}
	return nil;
}

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}

func DFS(r *[]int, used map[int]struct{}, n int) bool {
	if len(used) == n {
		return true
	}

	for i := 1; i <= n; i++ {
		if _, exists := used[i]; exists {
			continue
		}
		if len(*r) > 0 {
			lastElement := (*r)[len(*r)-1]
			if !isPerfectSquare(lastElement + i) {
				continue 
			}
		}

		used[i] = struct{}{}
		*r = append(*r, i)  

		if DFS(r, used, n) {
			return true
		}

		delete(used, i)      
		*r = (*r)[:len(*r)-1] 
	}
	return false
}