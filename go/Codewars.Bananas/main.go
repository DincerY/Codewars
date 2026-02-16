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
		for i, rn := range str {
			if rn == '-' {
				continue
			}
			if byte(rn) == banana[index] {
				index++
			} else {
				break
			}
			if index == len(banana) {
				if i == len(str)-1 {
					res = append(res, str)
				}
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
