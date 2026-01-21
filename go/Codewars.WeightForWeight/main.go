package main

import (
	"strings"
)

func main() {
	sumDigs("10")
	sumDigs("21")

	OrderWeight("103 123 4444 99 2000")
	OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123")

}

func sumDigs(value string) int {
	sum := 0
	for _, char := range value {
		sum += int(char) - 48
	}
	return sum
}

var dic map[int][]int = make(map[int][]int)

func OrderWeight(strng string) string {
	values := strings.Split(strng, " ")
	//len := len(values)

	for _, value := range values {
		print(value, sumDigs(value))
	}
	return ""
}
