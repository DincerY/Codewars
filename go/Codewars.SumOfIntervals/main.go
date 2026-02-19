package main

func main() {
	SumOfIntervals([][2]int{{1, 5}})
	SumOfIntervals([][2]int{{1, 5}, {6, 10}})
	SumOfIntervals([][2]int{{1, 5}, {1, 5}})
	SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})

}

func helper(intervals [][2]int, val int) int {
	index := -1
	for _, inter := range intervals {
		if inter[1] > val {
			return index
		}
	}
	return index
}

func SumOfIntervals(intervals [][2]int) int {
	arr := [][2]int{}
	res := 0

	for _, inter := range intervals {
		print(inter[0], inter[1])
		if len(arr) == 0 {
			arr = append(arr, inter)
		}

	}

	return res
}
