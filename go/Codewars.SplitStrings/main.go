package main

func main() {
	Solution2("abc")
	Solution2("dawsd")
	Solution2("awsaws")
}

func Solution(str string) []string {
	var res []string
	for i := 0; i < len(str); i += 2 {
		if i+2 > len(str) {
			b := str[i : i+1]
			b += "_"
			res = append(res, b)

		} else {
			a := str[i : i+2]
			res = append(res, a)

		}
	}
	return res
}

func Solution2(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
