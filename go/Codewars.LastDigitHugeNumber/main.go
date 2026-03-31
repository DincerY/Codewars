package main

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

func cPow(a, b, m int) int {
	if b == 0 {
		if 1 < m {
			return 1
		}
		return 1%m + m
	}

	if a == 0 {
		return 0
	}

	if a >= m {
		a = (a % m) + m
	}

	res := 1
	for i := 0; i < b; i++ {
		res *= a
		if res >= m {
			res = (res % m) + m
		}
	}
	return res
}

func modFor(i int) int {
	switch i {
	case 0:
		return 10
	case 1:
		return 4
	case 2:
		return 2
	default:
		return 1
	}
}

// LastDigit, kuvvet kulesinin son basamağını hesaplar
func LastDigit(as []int) int {
	n := len(as)

	if n == 0 {
		return 1
	}

	for i := 3; i < n; i++ {
		if as[i] > 0 {
			n = i + 1
			break
		}
	}

	exp := 1
	for i := n - 1; i >= 0; i-- {
		exp = cPow(as[i], exp, modFor(i))
	}

	return exp % 10
}
