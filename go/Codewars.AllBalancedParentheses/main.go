package main

func main() {
	BalancedParens(0)
	BalancedParens(1)
	BalancedParens(2)
	BalancedParens(3)
	BalancedParens(4)
}

var result []string

func BalancedParens(n int) []string {
	if n == 0 {
		return []string{""}
	}
	Recursive("", n, n, n*2)
	return result
}

func isValid(res string) bool {
	var stack []rune
	for _, chr := range res {
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

func Recursive(res string, left int, right int, resLen int) {
	if left > 0 {
		Recursive(res+"(", left-1, right, resLen)
	}
	if right > 0 {
		Recursive(res+")", left, right-1, resLen)
	}
	if left == 0 && right == 0 {
		if len(res) == resLen && isValid(res) {
			result = append(result, res)
		}
	}
}
