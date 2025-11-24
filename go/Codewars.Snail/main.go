package main

func main() {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Snail(snailMap)
	snailMap = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	Snail(snailMap)
}

func Snail(snaipMap [][]int) []int {
	row := len(snaipMap)
	column := len(snaipMap[0])

	println(column)
	println(row)

	return []int{}
}

func Recursive() {

	Recursive() // right
	Recursive() // down
	Recursive() // left
	Recursive() // up

}
