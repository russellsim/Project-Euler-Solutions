package main

import (
	"fmt"
	"time"

	"github.com/russellsim/Project-Euler-Solutions/lib"
)

func main() {

	start := time.Now()

	const limit = 1000000
	sieve := lib.NewPrimeSieve(limit)

	primesFound := 0
	currentNumber := 0

	for primesFound < 10001 {
		currentNumber++
		if sieve.IsPrime(currentNumber) {
			primesFound++
		}
	}

	fmt.Println("Result: ", currentNumber)
	fmt.Println("Time taken: ", time.Since(start))
}
