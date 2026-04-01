package main

func main() {
	DblLinear(10)
	DblLinear(50)
	DblLinear(100)

}

func DblLinear(n int) int {
	res := []int{1}
	dic := make(map[int]struct{})
	dic[1] = struct{}{}

	for true {
		i := 1
		if dic[i] == struct{}{} {
			continue
		}
		if i == n {
			break
		}
		i++
	}
	return res[len(res)-1]
}
