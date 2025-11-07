package main

func main() {
	var emptyArr []int
	ArrayDiff([]int{1, 2}, []int{1})
	ArrayDiff([]int{1, 2, 2}, []int{1})
	ArrayDiff([]int{1, 2, 2}, []int{2})
	ArrayDiff([]int{1, 2, 2}, emptyArr)
}

func contains(arr []int, val int) (bool, int) {
	for i, v := range arr {
		if v == val {
			return true, i
		}
	}
	return false, -1
}

func ArrayDiff(a, b []int) []int {
	//it is working one time. If you want to remove two times
	//It won't work.
	for i := 0; i < len(b); i++ {
		if contain, index := contains(a, b[i]); contain {
			a = append(a[:index], a[index+1:]...)
		}
	}

	return a
}