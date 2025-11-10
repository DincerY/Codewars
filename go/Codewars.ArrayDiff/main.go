package main

func main() {
	var emptyArr []int
	ArrayDiff([]int{1, 2}, []int{1})
	ArrayDiff([]int{1, 2, 2}, []int{1})
	ArrayDiff([]int{1, 2, 2}, []int{2})
	ArrayDiff([]int{1, 2, 2}, emptyArr)
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func ArrayDiff(a, b []int) []int {
	var res []int
	for i := 0; i < len(a); i++ {
		if !contains(b, a[i]) {
			res = append(res, a[i])
		}
	}
	return res
}
