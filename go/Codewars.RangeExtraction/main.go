package main

import "strconv"

func main() {
	Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
}

func Solution(list []int) string {
	res := ""
	first := list[0]
	last := list[0]
	for i := 1; i < len(list); i++ {
		if last == list[i]-1 {
			last = list[i]
		} else {
			if last-first >= 2 {
				res += strconv.Itoa(first) + "-" + strconv.Itoa(last)
			} else if last-first == 1 {
				res += strconv.Itoa(first) + "," + strconv.Itoa(last)
			} else {
				res += strconv.Itoa(first)
			}

			res += ","
			first = list[i]
			last = list[i]

		}
	}
	if last-first >= 2 {
		res += strconv.Itoa(first) + "-" + strconv.Itoa(last)
	} else if last-first == 1 {
		res += strconv.Itoa(first) + "," + strconv.Itoa(last)
	} else {
		res += strconv.Itoa(first)
	}

	return res
}
