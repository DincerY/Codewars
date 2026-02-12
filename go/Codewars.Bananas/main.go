package main

func main() {
	Bananas("bbananana")
	Bananas("banana")
	Bananas("bananaaa")
}

var allRecRes = []string{}

func Bananas(s string) []string {
	banana := "banana"
	res := []string{}
	Recursion(s, 0)
	for _, str := range allRecRes {
		index := 0
		for _, rn := range str {
			if byte(rn) == banana[index] {
				index++
			}
			if index == len(banana) {
				res = append(res, str)
				break
			}
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
