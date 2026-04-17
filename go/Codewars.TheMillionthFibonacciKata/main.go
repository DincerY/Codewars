package main

import (
	"math"
	"math/big"
)

func main() {
	fib(1)
	fib(2)
	fib(3)
	fib(5)
	fib(50)
}

func fib(n int64) *big.Int {
	if n < 0 {
		n = int64(math.Abs(float64(n)))
	}
	res := recursiondp(n)

	return big.NewInt(res)
}

var memo = make(map[int64]int64)

func recursiondp(n int64) int64 {
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 1 {
		return n
	}
	memo[n] = recursiondp(n-1) +
		recursiondp(n-2)
	return memo[n]
}
