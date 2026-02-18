package main

func main() {
	Bananas("bbananana")
	Bananas("banana")
	Bananas("bananaaa")
}

var allRecRes []string

func Bananas(s string) []string {
	allRecRes = []string{}
	banana := "banana"
	res := []string{}
	Recursion(s, 0)
	for test, str := range allRecRes {
		if test == 252 {
			print("")
		}
		index := 0
		indexForRn := -1
		for _, rn := range str {
			if rn == '-' {
				indexForRn++
				continue
			}
			if index >= len(banana) {
				break
			}
			if byte(rn) == banana[index] {
				indexForRn++
				index++
			} else {
				break
			}

		}
		if index == len(banana) && indexForRn == len(str)-1 {
			res = append(res, str)
		}
	}

	return res
}

func Recursion(s string, index int) {
	if index >= len(s) {
		allRecRes = append(allRecRes, s)
		return
	}
	prevS := s
	r := []rune(s)
	r[index] = '-'
	index++
	Recursion(string(r), index)
	Recursion(prevS, index)
}
