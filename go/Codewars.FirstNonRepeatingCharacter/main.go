package main

import "strings"

func main() {

	FirstNonRepeating("a")
	FirstNonRepeating("stress")
	FirstNonRepeating("sTress")

	FirstNonRepeating("moonmen")

	FirstNonRepeating("")
	FirstNonRepeating(",")

	FirstNonRepeating("abba")
	FirstNonRepeating("bb")

	FirstNonRepeating("~><#~><")
	FirstNonRepeating("hello world, eh?")
}

func FirstNonRepeating(str string) string {
	res := ""
	list := make(map[rune][2]int)
	if len(str) == 0 {
		return ""
	}

	lowerStr := strings.ToLower(str)

	for i, rn := range lowerStr {
		test := list[rn]
		if test[0] == 0 && test[1] == 0 {
			test[0] = i
			test[1] = 1
		} else {
			test[1]++
		}
		list[rn] = test
	}
	for _, lst := range list {
		if lst[0] == 1 {
			res = string(str[lst[1]])
		}
	}

	return res
}
