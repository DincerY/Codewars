package main

func main() {
	SelNumber(4000, 2)
	SelNumber(47, 3)
}

func SelNumber(n, d int) int {
	res := 0
	if n < 12 {
		return 0
	}
	for i := 3456; i <= n; i++ {
		j := i
		digit := []int{}
		for j > 0 {
			number := j % 10
			digit = append([]int{number}, digit...)
			j = j / 10
		}
		isValid := false
		for i := len(digit) - 1; i > 0; i-- {
			if digit[i]-digit[i-1] <= d && digit[i]-digit[i-1] > 0 && digit[i] != digit[i-1] {
				isValid = true
			} else {
				isValid = false
			}
		}
		if isValid {
			res++
		}
	}
	return res
}
