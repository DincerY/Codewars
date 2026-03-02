package main

func main() {
	MoveZeros2([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})
}

func helper(arr []int, firstIndex, secondIndex int) []int {
	firstVal := arr[firstIndex]
	arr[firstIndex] = arr[secondIndex]
	arr[secondIndex] = firstVal
	return arr
}

func MoveZeros(arr []int) []int {
	for index, val := range arr {
		if val == 0 {
			for i := index + 1; i < len(arr); i++ {
				if arr[i] != 0 {
					arr = helper(arr, index, i)
					break
				}
			}
		}
	}
	return arr
}

func MoveZeros2(arr []int) []int {
	result := make([]int, len(arr))
	index := 0
	for _, val := range arr {
		if val != 0 {
			result[index] = val
			index++
		}
	}
	return result
}
