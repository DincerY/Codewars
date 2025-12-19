package main

func main() {
	LCS("a", "b")                  //""
	LCS("abcdef", "abc")           //"abc"
	LCS("abcdef", "acf")           //"acf"
	LCS2("132535365", "123456789") //"12356"
}

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	var result string
	Recursive(s1, s2, 0, 0, "", &result)

	return result
}

func Recursive(s1, s2 string, s1Index, s2Index int, temp string, result *string) {
	if s1Index >= len(s1) || s2Index >= len(s2) {
		if len(temp) > len(*result) {
			*result = temp
		}
		return
	}
	if s1[s1Index] == s2[s2Index] {
		temp += string(s1[s1Index])
		Recursive(s1, s2, s1Index+1, s2Index+1, temp, result)
		temp = temp[:len(temp)-1]
	}
	Recursive(s1, s2, s1Index+1, s2Index, temp, result)
	Recursive(s1, s2, s1Index, s2Index+1, temp, result)
}

// It is not mine solution
func LCS2(s1, s2 string) string {
	if len(s1)*len(s2) == 0 {
		return ""
	}

	if s1[0] == s2[0] {
		return s1[:1] + LCS2(s1[1:], s2[1:])
	}

	w1, w2 := LCS2(s1, s2[1:]), LCS2(s1[1:], s2)
	if len(w2) > len(w1) {
		return w2
	}

	return w1
}
