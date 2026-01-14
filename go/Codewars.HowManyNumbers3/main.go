package main

func main() {
	Recursion(0, 0, 0, 3, 0, 10)

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

// I should change parameters
func Recursion(digitCount, prevVal, sumDig, digs, test, resDig int) {
	for i := 1; i < 10; i++ {
		if digitCount < digs {
			if test == 0 {
				test = 1
			} else {
				test = test*10 + i
			}
			if sumDig == resDig {
				//result+1
			}
			Recursion(digitCount+1, i, sumDig+i, digs, test, resDig)
			test /= 10
		} else {
			break
		}
	}
}
