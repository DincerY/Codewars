package main

func main() {
	Permutations("abc")
	Permutations("a")
	Permutations("ab")
	Permutations("aabb")
	Permutations("abcd")
}

// working on res will get understading solution
var res []string

func Permutations(s string) []string {
	Recursion(s)
	return nil
}

func Recursion(str string) {
	for i := 0; i < len(str); i++ {
		res = append(res, string(str))
		Recursion(str[:i] + str[i+1:])
		Recursion(str[:i] + str[i+1:])
	}

}
