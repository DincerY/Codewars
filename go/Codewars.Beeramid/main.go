package main

func main() {
	Beeramid(-1, 4)
	Beeramid(2, 1.9)

	Beeramid(21, 1.5)
	Beeramid(1500, 2)

	Beeramid(9, 2)
	Beeramid(9, 1.5)

}

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 || price <= 0 {
		return 0
	}
	levels := 0

	can := int(float64(bonus) / price)

	for i := 1; i <= can; i++ {
		square := i * i
		if square <= can {
			levels++
			can -= square
		} else {
			break
		}
	}

	return levels
}
