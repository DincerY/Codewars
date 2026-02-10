package main

func main() {
	CreateSpiral(5)

	CreateSpiral(1)
	CreateSpiral(2)
	CreateSpiral(3)
	CreateSpiral(4)
}

func CreateSpiral(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	left, top := 0, 0
	right, bottom := n, n

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	lastAddedValue := 1

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			res[top][i] = lastAddedValue
			lastAddedValue++
		}
		top++

		for i := top; i < bottom; i++ {
			res[i][right-1] = lastAddedValue
			lastAddedValue++
		}
		right--

		for i := right - 1; i >= left; i-- {
			res[bottom-1][i] = lastAddedValue
			lastAddedValue++
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			res[i][left] = lastAddedValue
			lastAddedValue++
		}
		left++

	}
	return res
}
