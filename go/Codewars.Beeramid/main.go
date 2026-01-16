package main

func main() {
	Beeramid(9, 2)
	Beeramid(9, 1.5)
	Beeramid(-1, 4)

}

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 {
		return 0
	}
	levels := 1

	can := int(float64(bonus) / price)
	remainder := can

	for i := 1; i < remainder; i++ {
		levels++
		square := i * i
		if square < remainder {
			remainder -= square
		}
	}

	return levels
}
