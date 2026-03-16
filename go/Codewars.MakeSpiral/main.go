package main

func main() {
	Spiralize(5)
	Spiralize(8)

	Spiralize(10)
}

func Spiralize(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
	}
	left, top := 0, 0
	right, bottom := size-1, size-1
	for left <= right && top <= bottom {
		if left > 0 {
			res[top][left-1] = 1
		}
		for i := left; i <= right; i++ {
			res[top][i] = 1
		}
		for i := top; i <= bottom; i++ {
			res[i][right] = 1
		}
		top += 2
		for i := right; i >= left; i-- {
			res[bottom][i] = 1
		}
		right -= 2
		for i := bottom; i >= top; i-- {
			res[i][left] = 1
		}
		bottom -= 2
		left += 2
	}
	return res
}
