package main

func main() {
	CreateSpiral(5)

	CreateSpiral(1)
	CreateSpiral(2)
	CreateSpiral(3)
	CreateSpiral(4)
}

func CreateSpiral(n int) [][]int {

	var left, right, top, bottom int
	left = 0
	right = n
	top = 0
	bottom = n

	//test matrix
	res := [][]int{}
	/*subRes := &res[0]
	*subRes = append(*subRes, 1)
	res = append(res, []int{1, 2, 3, 4, 5, 6})*/

	for i := 0; i < n; i++ {
		res = append(res, []int{})
		for j := 0; j < n; j++ {
			res[i] = append(res[i], 0)
		}
	}
	// your code here

	for left != right && top != bottom {
		lastAddedValue := 1
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
