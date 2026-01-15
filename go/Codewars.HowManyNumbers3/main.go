package main

func main() {
	FindAll(84, 4)
	FindAll(10, 3)
	FindAll(27, 3)
}

var result []int

func FindAll(sumDig, digs int) []int {
	result = []int{}
	Recursion(1, 0, digs, sumDig, -1)
	if len(result) == 0 {
		return []int{}
	}
	return []int{len(result), result[0], result[len(result)-1]}
}

func Recursion(lastDigit, digitCount, expectedDigs, sumDig, test int) {
	for i := lastDigit; i < 10; i++ {
		if digitCount < expectedDigs {
			if test == -1 {
				test = 1
			} else {
				test = test*10 + i
			}

			Recursion(i, digitCount+1, expectedDigs, sumDig-i, test)
			test /= 10
		} else if digitCount == expectedDigs {
			if sumDig == 0 {
				result = append(result, test)
			}
			break
		} else {
			break
		}
	}
}
