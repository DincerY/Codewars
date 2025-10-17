package main

func main() {
	Multiple3And5(100)
}

func Multiple3And5(number int) int {
	var res int = 0
	for i := 0; i < number; i++ {
		if i%3 == 0 {
			res += i
		} else if i%5 == 0 {
			res += i
		} else {
			continue
		}
	}
	return res
}
func Multiple3And5Dif(number int) int {
	var res int = 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}
	return res
}