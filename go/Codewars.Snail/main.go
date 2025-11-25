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
	if len(snaipMap) == 0 {
		return nil
	}
	row := len(snaipMap)
	column := len(snaipMap[0])

	var res []int
	Recursive(0, 0, row-1, column-1, &res, snaipMap)

	return res
}

// helper func
func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func Recursive(x, y, row, col int, res *[]int, snailMap [][]int) {

	if x > row || x < 0 || y > col || y < 0 || contains(*res, snailMap[x][y]) {
		return
	} else {
		*res = append(*res, snailMap[x][y])
	}
	if x-1 < 0 || contains(*res, snailMap[x-1][y]) {
		Recursive(x, y+1, row, col, res, snailMap) // right
	}
	Recursive(x+1, y, row, col, res, snailMap) // down
	Recursive(x, y-1, row, col, res, snailMap) // left
	Recursive(x-1, y, row, col, res, snailMap) // up
}
