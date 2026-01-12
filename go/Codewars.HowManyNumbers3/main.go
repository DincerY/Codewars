package main

func main() {

}

func FindAll(sumDig, digs int) []int {
	res := []int{}

	//this code only use with 3 digits. I will solve it
	for i := 1; i < 9; i++ {
		total := 0
		total += i
		for j := i; j < 9; j++ {
			total += j
			for k := j; k < 9; k++ {
				total += k
				if total == sumDig {

				}
			}
		}
	}

	//return [amount, min, max]
	return []int{len(res), res[0], res[len(res)-1]}
}
