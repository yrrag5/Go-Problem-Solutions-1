// Author: Garry Cummins
// Date: 06/10/2017

// Adapted from: https://play.golang.org/p/5FUOzjBa-o

package main

import (
 "fmt"
 "strings"
)

func main() {

	var v string
	//Takes in user input 
	fmt.Println("Enter a string value:")
	fmt.Scanf("%s\n", &v)

	v = strings.ToLower(v)
	fmt.Println(isPalindrome(v))
}

// Function that checks to see if the user input is or is not 
// a palimdrome 
func isPalindrome(s string) string {

	// Checks middle and last letter of the inputted word
	m := len(s) / 2
	l := len(s) - 1

	// Compares the first and last letter 
	for i := 0; i < m; i++ {

		if s[i] != s[l-i] {
			return "That is not a Palimdrome!"
		}
	}
	return "That is a Palimdrome!"
}