package main

import "fmt"

func main() {
	Decomp(22)

	Decomp(5)
	Decomp(12)
	Decomp(25)
}

func helper(val, divisor int) (int, int) {
	res := 0
	for val%divisor == 0 {
		val /= divisor
		res++
	}
	return res, val
}

func Decomp(n int) string {
	var res string
	arr := []int{}
	for i := 2; i <= n; i++ {
		arr = append(arr, i)
	}

	resMap := map[int]int{}

	for len(arr) > 0 {

		for len(arr) > 0 && arr[0] == 1 {
			arr = arr[1:]
		}

		if len(arr) <= 0 {
			break
		}

		prime := arr[0]
		for i := 0; i < len(arr); i++ {
			if arr[i] == 1 {
				continue
			}
			power, remainder := helper(arr[i], prime)
			resMap[prime] += power
			arr[i] = remainder
		}
	}
	for i := 2; i <= n; i++ {
		if len(resMap) <= 0 {
			break
		}
		value := resMap[i]
		if value == 0 {
			continue
		}
		if value == 1 {
			res += fmt.Sprintf("%d", i)
		} else {
			res += fmt.Sprintf("%d^%d", i, value)
		}
		res += " * "
	}
	return res
}
