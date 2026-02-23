package main

func main() {
	SumOfIntervals([][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}})
	SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})

	SumOfIntervals([][2]int{{1, 5}})
	SumOfIntervals([][2]int{{1, 5}, {6, 10}})
	SumOfIntervals([][2]int{{1, 5}, {1, 5}})

}

func helper(intervals [][2]int, interval [2]int) []int {
	res := []int{}
	for index, inter := range intervals {
		if inter[0] < interval[0] && interval[0] < inter[1] && interval[1] > inter[1] {
			res = append(res, index)
		}
	}
	return res
}

// func helper(intervals [][2]int, val int) []int {
// 	res := []int{}
// 	for index, inter := range intervals {
// 		if inter[0] < val && inter[1] > val {
// 			res = append(res, index)
// 		}
// 	}
// 	return res
// }

func SumOfIntervals(intervals [][2]int) int {
	res := 0

	for intervalsIndex, inter := range intervals {
		index := helper(intervals, inter)
		for _, i := range index {
			intervals[i][1] = inter[1]
			intervals = append(intervals[:intervalsIndex], intervals[intervalsIndex+1:]...)
		}

	}

	for _, inter := range intervals {
		res += inter[1] - inter[0]

	}
	return res
}

// func SumOfIntervals2(intervals [][2]int) int {
// 	arr := [][2]int{}
// 	res := 0

// 	for _, inter := range intervals {
// 		print(inter[0], inter[1])
// 		if len(arr) == 0 {
// 			arr = append(arr, inter)
// 		}

// 	}

// 	return res
// }
