package main

func main() {
	DblLinear(1000)

	DblLinear(10)
	DblLinear(50)
	DblLinear(100)

}

func DblLinear(n int) int {
	res := make([]int, n+1)
	res[0] = 1

	i, j := 0, 0

	for k := 1; k <= n; k++ {
		twoRes := 2*res[i] + 1
		threeRes := 3*res[j] + 1

		if twoRes < threeRes {
			res[k] = twoRes
			i++
		} else if twoRes > threeRes {
			res[k] = threeRes
			j++
		} else {
			res[k] = threeRes
			i++
			j++
		}
	}
	return res[n]
}
