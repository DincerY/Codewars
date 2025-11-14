package main

func main() {
	PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3})
	PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1})
	PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2, 1})

}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

type IsPeak struct {
	Left  bool
	Right bool
	Val   int
}

func PickPeaks(array []int) PosPeaks {
	res := map[int]*IsPeak{}

	var index int
	for i := 1; i < len(array)-1; i++ {
		if array[i] != array[i-1] {
			index = i
		}
		res[i] = &IsPeak{
			Val: array[i],
		}
		if array[i] > array[i-1] {
			res[index].Left = true
		}
		if array[i] > array[i+1] {
			res[index].Right = true
		}

	}
	pos := PosPeaks{}

	for i := 1; i < len(array)-1; i++ {
		if res[i].Left && res[i].Right {
			pos.Pos = append(pos.Pos, i)
			pos.Peaks = append(pos.Peaks, res[i].Val)
		}
	}
	return pos
}
