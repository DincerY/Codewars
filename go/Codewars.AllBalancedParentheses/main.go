package main

func main() {
	BalancedParens(4)

	BalancedParens(0)
	BalancedParens(1)
	BalancedParens(2)
	BalancedParens(3)

}

func BalancedParens(n int) []string {
	Recursive("", n, n)
	return nil
}

var test []string

func Recursive(res string, left int, right int) {

	if left > 0 {
		Recursive(res+"(", left-1, right)
		left++
	}
	if right > 0 {
		Recursive(res+")", left, right-1)
		right++
	}
	if left == 0 && right == 0 {
		test = append(test, res)
	}
}
