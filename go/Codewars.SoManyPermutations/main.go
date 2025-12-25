package main

func main() {
	Permutations("aa")

	Permutations("ab")

	Permutations("abc")
	Permutations("a")
	Permutations("aabb")
	Permutations("abcd")
}

var resMap map[string]struct{}

func Permutations(s string) []string {
	resMap = make(map[string]struct{})
	Recursion(s, "", len(s))
	var result []string
	for k := range resMap {
		result = append(result, k)
	}
	return result
}

func Recursion(str string, temp string, lenght int) {
	if len(temp) == lenght {
		resMap[temp] = struct{}{}
		return
	}
	for i := 0; i < len(str); i++ {
		Recursion(str[:i]+str[i+1:], temp+string(str[i]), lenght)
		Recursion(str[:i]+str[i+1:], temp, lenght)
	}
}

// It is not mine
func Permutations2(s string) (a []string) {
	if len(s) < 2 {
		return []string{s}
	}
	m := map[string]bool{}
	for _, sub := range Permutations(s[1:]) {
		for i := 0; i <= len(sub); i++ {
			st := sub[0:i] + s[0:1] + sub[i:]
			if m[st] {
				continue
			}
			m[st] = true
			a = append(a, st)
		}
	}
	return
}
