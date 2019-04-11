package main

import (
	"fmt"
	"math"
	"time"

	"github.com/russellsim/Project-Euler-Solutions/lib"
)

func main() {

	start := time.Now()

	const number = 600851475143
	limit := int(math.Ceil(math.Sqrt(number)))
	sieve := lib.NewPrimeSieve(limit)

	sieveTime := time.Since(start)

	for i := limit; i >= 2; i-- {
		if number%i == 0 {
			if sieve.IsPrime(i) {
				fmt.Println("Result is", i)
				break
			}
		}
	}

	fmt.Println("Time taken to setup sieve", sieveTime)
	fmt.Println("Time taken to compute", time.Since(start))
}
