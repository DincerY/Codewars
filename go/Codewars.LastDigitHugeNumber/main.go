package main

import "math"

func main() {
	LastDigit([]int{937640, 767456, 981242})

	LastDigit([]int{})
	LastDigit([]int{0, 0})
	LastDigit([]int{0, 0, 0})
	LastDigit([]int{1, 2})
	LastDigit([]int{3, 4, 5})
	LastDigit([]int{4, 3, 6})

	LastDigit([]int{3, 4, 2})
	LastDigit([]int{9, 9, 9})

}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func LastDigit(arr []int) int {
	if len(arr) == 0 {
		return 1
	}

	exp := 1

	for i := len(arr) - 1; i >= 1; i-- {
		base := arr[i]

		if exp == 0 {
			exp = 1
		}

		exp = modPow(base, exp, 4)

		if exp == 0 {
			exp = 4
		}
	}

	return modPow(arr[0], exp, 10)
}

func LastDigit2(as []int) int {
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

	for i := 0; i < power-1; i++ {
		lastDigit = (lastDigit * val) % 10

	}
	return lastDigit
}
