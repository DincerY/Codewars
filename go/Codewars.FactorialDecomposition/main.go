package main

func main() {
	Decomp(5)
	Decomp(12)
	Decomp(22)
	Decomp(25)
}

func helper(val, divisor int) int {
	res := 0
	for val%divisor == 0 {
		val /= divisor
		res++
	}
	return res
}

func Decomp(n int) string {
	var res string
	arr := []int{}
	for i := 2; i <= n; i++ {
		arr = append(arr, i)
	}
	powerCount := 0
	value := arr[0]

	for len(arr) > 0 {

	}

	res += string(value)
	return ""
}
