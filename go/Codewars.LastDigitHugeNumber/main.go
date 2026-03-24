package main

import "math"

func main() {
	LastDigit([]int{})
	LastDigit([]int{0, 0})
	LastDigit([]int{0, 0, 0})
	LastDigit([]int{1, 2})
	LastDigit([]int{3, 4, 5}) //1
	LastDigit([]int{4, 3, 6})

	LastDigit([]int{3, 4, 2})
	LastDigit([]int{9, 9, 9})

}

func LastDigit(as []int) int {
	if len(as) == 0 {
		return 1
	}
	if len(as) == 1 {
		return as[0] % 10
	}

	floatAs := []float64{}

	for _, val := range as {
		floatAs = append(floatAs, float64(val))
	}

	for i := len(floatAs) - 2; i > 0; i-- {
		floatAs[i] = math.Pow(floatAs[i], floatAs[i+1])
	}

	power := int(floatAs[1])
	val := int(floatAs[0])
	lastDigit := val % 10
	if power == 0 {
		return 1
	}

	for i := 0; i < power; i++ {
		lastDigit = (lastDigit * val) % 10

	}
	return lastDigit
}
