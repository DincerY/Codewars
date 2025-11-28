package main

func main() {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	/*Snail(snailMap)
	snailMap = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	Snail(snailMap)*/

	snailMap = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	Snail(snailMap)
}

func Snail(snaipMap [][]int) []int {
	var res []int
	row := len(snaipMap)
	column := len(snaipMap[0])
	left := 0
	top := 0
	right := row
	bottom := column
	for left != top && right != bottom {
		for i := left; i < right; i++ {
			res = append(res, snaipMap[top][i])
		}
		top++
		for i := top; i < right; i++ {
			res = append(res, snaipMap[i][right-1])
		}
		right--
		for i := right - 1; i >= left; i-- {
			res = append(res, snaipMap[bottom-1][i])
		}
		bottom--
		for i := bottom - 1; i >= top; i-- {
			res = append(res, snaipMap[i][left])
		}
		left++
	}
	return res
}
