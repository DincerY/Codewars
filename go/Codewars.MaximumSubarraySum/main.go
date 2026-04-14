package main

func main() {
	//MaximumSubarraySum2([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})

	MaximumSubarraySum2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	MaximumSubarraySum([]int{1, 2, 3, 4, 5})

}

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	res := -21000000
	max := -21000000

	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > max {
				max = sum
			}
		}
		if max > res {
			res = max
		}
	}
	if res < 0 {
		return 0
	} else {
		return res
	}
}

func MaximumSubarraySum2(numbers []int) int {
	res, sum := 0, 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}
