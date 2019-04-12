package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	//divisors := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Start from the back and remove duplicates.
	// 20 is divisible by 2, 4, 5 and 10, so remove them
	// 18 is divisible by 3, 6, so remove them
	// 16 is divisble by 8, so remove it
	// 14 is divisble by 7, so remove it
	divisors := []int{9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	result := 20
	found := false
	for !found {
		result += 20 // Simple optimization, know that number must be a multiple of 20
		found = true
		for _, div := range divisors {
			if result%div != 0 {
				found = false
				break
			}
		}
	}

	fmt.Println("Result:", result)
	fmt.Println("Time taken: ", time.Since(start))
}
