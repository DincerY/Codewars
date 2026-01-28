package main

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
	//powerCount := 0
	value := arr[0]
	resMap := map[int]int{}

	for len(arr) > 0 {
		for arr[0] == 1 {
			arr = arr[1:]
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

	res += string(value)
	return ""
}
