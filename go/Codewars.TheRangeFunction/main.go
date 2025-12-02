package main

import "math"

func main() {
	Range3(-40, 5, 4)

	Range3(1, 4, 0)

	Range(0, 10, 1)
	Range(1, 11, 1)
	Range(0, 30, 5)

	Range(10, 0, 0)

}

func Range(stard, end, stop int) []int {
	if stard > end {
		return nil
	}
	res := []int{}
	step := 0
	if stop == 0 {
		step = (end - stard) / 1
	} else {
		step = int(math.Ceil(float64((end - stard)) / float64(stop)))
	}

	for i := stard; i < end; i += stop {
		if step == 0 {
			break
		}
		res = append(res, i)
		step--
	}
	return res
}

func Range2(stard, end, stop int) []int {
	if stard > end {
		return nil
	}
	res := []int{}
	value := stard
	step := stop
	if stop == 0 {
		step = 1
	}

	for i := stard; i < end; i += stop {
		if value == end {
			break
		}
		res = append(res, i)
		value += step
	}
	return res
}

func Range3(stard, end, stop int) (res []int) {
	for i := stard; i < end && stard < end; stard += stop {
		res = append(res, stard)
		i++
	}
	return
}
