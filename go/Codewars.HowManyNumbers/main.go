package main

func main() {
	SelNumber(20, 2)
	SelNumber(47, 3)
}

func SelNumber(n, d int) int {
	res := 0
	if n < 12 {
		return 0
	}
	for i := 12; i <= n; i++ {
		j := i
		var digit []int
		for j > 0 {
			number := j % 10
			digit = append([]int{number}, digit...)
			j = j / 10
		}
		if true {
			res++
		}
	}
	return res
}
