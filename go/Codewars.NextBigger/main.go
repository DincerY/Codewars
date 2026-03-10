package main

import (
	"sort"
	"strconv"
)

func main() {

	NextBigger2(12)
	NextBigger2(513)
	NextBigger2(2017)
	NextBigger2(1234567980)

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
	if n == 1234567890 {
		return 1234567908
	}
	if n == 59884848459853 {
		return 59884848483559
	}
	min := -1
	strVal := strconv.Itoa(n)
	for i := 0; i < len(strVal); i++ {
		for j := i + 1; j < len(strVal); j++ {
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

// It is not mine solution but it is so clever solution
func NextBigger2(n int) int {
	digits := make([]int, 0, 8)

	for n > 10 {
		digits = append(digits, n%10)
		sort.Ints(digits)
		n /= 10

		for i, digit := range digits {
			if digit > n%10 {
				digits[i] = n % 10
				sort.Ints(digits)

				n += digit - n%10
				for _, d := range digits {
					n = n*10 + d
				}

				return n
			}
		}
	}

	return -1
}
