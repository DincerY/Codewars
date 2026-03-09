package main

import "strconv"

func main() {
	NextBigger(12)
	NextBigger(513)
	NextBigger(2017)
	NextBigger(414)
	NextBigger(144)
}

func helper(l, r int, val string) int {
	arr := []rune(val)
	arr[l], arr[r] = arr[r], arr[l]
	intRes, _ := strconv.Atoi(string(arr))
	return intRes
}

func NextBigger(n int) int {
	min := -1
	strVal := strconv.Itoa(n)
	for i := 0; i < len(strVal); i++ {
		for j := i; j < len(strVal); j++ {
			if i == j {
				continue
			}
			newVal := helper(i, j, strVal)
			if newVal > n {
				if min == -1 {
					min = newVal
				} else {
					if min > newVal {
						min = newVal
					}
				}
			}
		}
	}

	return min
}
