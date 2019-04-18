package main

import (
	"fmt"
	"time"

	"github.com/russellsim/Project-Euler-Solutions/lib"
)

func main() {
	start := time.Now()

	limit := 2000000
	sieve := lib.NewPrimeSieve(limit)

	sieveTime := time.Since(start)

	result := 0

	for i := 2; i < limit; i++ {
		if sieve.IsPrime(i) {
			result += i
		}
	}

	fmt.Println("Result:", result)
	fmt.Println("Time taken to setup sieve:", sieveTime)
	fmt.Println("Time taken to compute:", time.Since(start))
}
