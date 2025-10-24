package main

func main() {
	SquareSumsRow(10)
}

func SquareSumsRow(n int) []int {
	var square []int = []int{4, 9, 16, 25, 36, 49, 64, 81, 100}
	var test = [][]int{}

	for i := 0; i <= n; i++ {
		test = append(test, []int{})
	}
	for i := 1; i <= n; i++ {
		for _, val := range square {
			if val-i <= n {
				test[i] = append(test[i], 1)
			}
		}
	}

	return nil
}
