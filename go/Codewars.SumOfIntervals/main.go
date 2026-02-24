package main

func main() {
	SumOfIntervals([][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}})
	SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})

	SumOfIntervals([][2]int{{1, 5}})
	SumOfIntervals([][2]int{{1, 5}, {6, 10}})
	SumOfIntervals([][2]int{{1, 5}, {1, 5}})

}

func SumOfIntervals(intervals [][2]int) int {
	res := 0
	list := [][2]int{}
	for _, interval := range intervals {
		if len(list) == 0 {
			list = append(list, interval)
			continue
		}
		for _, item := range list {
			//0. elems must be updated
			if interval[0] < item[0] {
				item[0] = interval[0]
				//1. elems must be updated
			} else if item[1] > interval[0] && interval[1] > item[1] {
				item[1] = interval[1]
			}
		}
	}

	for _, item := range list {
		res += item[1] - item[0]
	}
	return res
}
