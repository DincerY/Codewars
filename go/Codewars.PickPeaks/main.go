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
	//It is working but if we have consecutive peaks it will not working
	res := map[int]*IsPeak{}

	for i := 1; i < len(array)-1; i++ {
		res[i] = &IsPeak{
			Val: array[i],
		}
		if array[i] > array[i-1] {
			res[i].Left = true
		}
		if array[i] > array[i+1] {
			res[i].Right = true
		}
	}
	pos := PosPeaks{}
	for a, b := range res {
		if b.Left && b.Right {
			pos.Pos = append(pos.Pos, a)
			pos.Peaks = append(pos.Peaks, b.Val)
		}
	}
	return pos
}
