package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	startRec := time.Now()
	res := recursiondp(50)
	durationRec := time.Since(startRec)
	fmt.Printf("DP Sonuç: %d | Süre: %v\n", res, durationRec)

	startRec = time.Now()
	res = recursion(50)
	durationRec = time.Since(startRec)
	fmt.Printf("Recursive Sonuç: %d | Süre: %v\n", res, durationRec)
	fib(1)
	fib(2)
	fib(3)
	fib(5)
	fib(50)
}

var memo = make(map[int]int)

func fib(n int64) *big.Int {
	return big.NewInt(n)
}

func recursion(n int) int {
	if n <= 1 {
		return n
	}
	return recursion(n-1) +
		recursion(n-2)
}

func recursiondp(n int) int {
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
