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

func Hammer(n int) uint {
	valueTable := []uint{1}

	power := []uint{2, 3, 5}
	res := map[uint]struct{}{}
	for len(res) != n {
		value := valueTable[0]
		res[value] = struct{}{}
		for _, p := range power {
			valueTable = append(valueTable, value*p)
		}
		valueTable = valueTable[1:]
	}
	//for res
	return 0
}
