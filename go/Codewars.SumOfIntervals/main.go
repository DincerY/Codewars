package main

import (
	"sort"
)

func main() {
	SumOfIntervals([][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}})
	SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})

	SumOfIntervals([][2]int{{1, 5}})
	SumOfIntervals([][2]int{{1, 5}, {6, 10}})
	SumOfIntervals([][2]int{{1, 5}, {1, 5}})

}

func SumOfIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	totalLength := 0
	currentStart := intervals[0][0]
	currentEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		nextStart := intervals[i][0]
		nextEnd := intervals[i][1]

		if nextStart < currentEnd {
			if nextEnd > currentEnd {
				currentEnd = nextEnd
			}
		} else {
			totalLength += currentEnd - currentStart

			currentStart = nextStart
			currentEnd = nextEnd
		}
	}
	totalLength += currentEnd - currentStart
	return totalLength
}
