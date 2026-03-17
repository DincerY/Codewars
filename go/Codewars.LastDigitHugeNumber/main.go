package main

import "math"

func main() {
	LastDigit([]int{3, 4, 2})
	LastDigit([]int{9, 9, 9})

}

func LastDigit(as []int) int {
	floatAs := []float64{}

	for _, val := range as {
		floatAs = append(floatAs, float64(val))
	}

	for i := len(floatAs) - 2; i >= 0; i-- {
		floatAs[i] = math.Pow(floatAs[i], floatAs[i+1])
	}
	return int(floatAs[0]) % 10
}
