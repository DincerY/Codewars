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
	right := 0
	bottom := right + 1
	left := column - 1
	top := row - 1
	for true {
		for right < column {
			res = append(res, snaipMap[top][right])
			right++
		}
		for bottom < row {
			res = append(res, snaipMap[bottom][right-1])
			bottom++
		}
		for left >= 0 {
			res = append(res, snaipMap[bottom-1][left])
			left--
		}
		for top >= 0 {
			res = append(res, snaipMap[top][left+1])
			top--
		}

	}
	return res
}
