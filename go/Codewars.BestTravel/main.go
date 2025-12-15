package main

import "math"

func main() {
	var ts = []int{50, 55, 56, 57, 58}
	ChooseBestSum(163, 3, ts)
	ts = []int{50}
	ChooseBestSum(163, 3, ts)
}

func ChooseBestSum(t, k int, ls []int) int {
	var result = 0
	Recursion(0, ls, k, 0, t, &result)
	if result == 0 {
		return -1
	}
	return result
}

func Recursion(start int, array []int, town int, sum int, expect int, result *int) {
	if *result == expect {
		return
	}
	if town == 0 {
		if sum <= expect {
			*result = int(math.Max(float64(*result), float64(sum)))
		}
		return
	}
	for i := start; i < len(array); i++ {
		Recursion(i+1, array, town-1, sum+array[i], expect, result)
	}
}
