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
	values := strings.Fields(strng)

	for _, value := range values {
		dic[sumDigs(value)] = append(dic[sumDigs(value)], value)
	}

	for len(dic) > 0 {
		minKey := 21000000
		for key := range dic {
			if key < minKey {
				minKey = key
			}
		}
		values := dic[minKey]
		if len(values) > 1 {
			for len(values) > 0 {
				minStr := "99999999999"
				minIndex := 0
				for index, str := range values {
					if str < minStr {
						minStr = str
						minIndex = index
					}
				}
				values = append(values[:minIndex], values[minIndex+1:]...)
				res += " " + minStr
			}
		} else {
			res += " " + values[0]
		}
		delete(dic, minKey)
	}
	if len(res) == 0 {
		return res
	}
	return res[1:]
}
