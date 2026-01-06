package main

func main() {
	Hammer(10)
	Hammer(1)
	Hammer(2)
	Hammer(3)
	Hammer(4)
	Hammer(5)
	Hammer(6)
	Hammer(7)
	Hammer(8)
	Hammer(9)
	Hammer(11)
	Hammer(12)
	Hammer(13)
	Hammer(14)
	Hammer(15)
}

func helper(values []uint) (uint, int) {
	min := values[0]
	index := 0
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
			index = i
		}
	}
	return min, index
}

func Hammer(n int) uint {
	valueTable := []uint{1}

	power := []uint{2, 3, 5}
	res := map[uint]struct{}{}
	for len(res) != n {
		value, index := helper(valueTable)
		res[value] = struct{}{}
		for _, p := range power {
			valueTable = append(valueTable, value*p)
		}
		valueTable = append(valueTable[:index], valueTable[index+1:]...)
	}
	var max uint = 0
	for k := range res {
		if k > max {
			max = k
		}
	}
	return result[len(result)-1]
}
