package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	palindromes := []int{1, 11, 101, 1001, 12321}
	notPalindromes := []int{13, 12341}

	// Check for palindromes
	for i := 0; i < len(palindromes); i++ {
		if isPalindrome(palindromes[i]) != true {
			t.Error(i, " is supposed to be a palindrome")
		}
	}

	// Check for not palindromes
	for i := 0; i < len(notPalindromes); i++ {
		if isPalindrome(notPalindromes[i]) != false {
			t.Error(i, " is not supposed to be a palindrome")
		}
	}
}
