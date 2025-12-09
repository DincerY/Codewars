package main

func main() {
	BalancedParens(0)
	BalancedParens(1)
	BalancedParens(2)
	BalancedParens(3)
	BalancedParens(4)
}

func BalancedParens(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var result []string

	Recursive("", n, n, &result)
	return result
}

func isValid(str string) bool {
	var stack []rune
	for _, chr := range str {
		if chr == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, chr)
		}
	}
	return len(stack) == 0
}

func Recursive(res string, left int, right int, result *[]string) {
	if left > 0 {
		Recursive(res+"(", left-1, right, result)
	}
	if right > 0 {
		Recursive(res+")", left, right-1, result)
	}
	if left == 0 && right == 0 {
		if isValid(res) {
			*result = append(*result, res)
		}
	}
}
