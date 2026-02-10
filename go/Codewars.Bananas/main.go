package main

func main() {
	Bananas("banana")
	Bananas("bbananana")
	Bananas("bananaaa")
}

func Bananas(s string) []string {

	return []string{}
}

func Recursion(s string, index int) {
	prevS := s
	r := []rune(s)
	r[index] = '-'
	index++
	Recursion(string(r), index)
	Recursion(prevS, index)
}
