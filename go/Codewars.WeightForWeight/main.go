package main

import (
	"strings"
)

func main() {
	sumDigs("10")
	sumDigs("21")

	OrderWeight("56 65 74 100 99 68 86 180 90")
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

var dic map[int][]string = make(map[int][]string)

func OrderWeight(strng string) string {
	var res string
	values := strings.Split(strng, " ")
	//len := len(values)

	//while append is being done, I should sort array
	for _, value := range values {
		dic[sumDigs(value)] = append(dic[sumDigs(value)], value)
		print(value, sumDigs(value))
	}

	for i := 0; i < len(dic); i++ {
		minKey := 21000000
		for key := range dic {
			if key < minKey {
				minKey = key
			}
		}
		values := dic[minKey]
		if len(values) > 1 {
			for i := 0; i < len(values); i++ {
				minStr := values[i]
				for _, str := range values {
					if str[0] < minStr[0] {
						minStr = str
					}
				}
				values = values[1:len(values)]
				res += minStr
			}
		} else {
			res += values[0]
		}
		delete(dic, minKey)
	}
	return res
}
