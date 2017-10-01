  
// Author: Garry Cummins
// Date: 01-10-2017

// Adapted from : http://wiki.c2.com/?FizzBuzzTest

package main

import "fmt"

func main() {

//Loops from 1 to 100
 for i := 1; i <= 100; i++ {

				//If the number is a multiple of 3, prints out Fizz
        if i%3 == 0 {
            fmt.Printf("Fizz")
        }

				//If the number is a multiple of 5, prints out Buzz
        if i%5 == 0 {
            fmt.Printf("Buzz")
        }

				//Will only print out number if the number is not a multiple of 3 and 5
        if i%3 != 0 && i%5 != 0 {
					  fmt.Printf("%d", i)
        }

        // When the number is a multiple of 3 and 5, will have both Fizz 
				//and Buzz print out on the same line 
        fmt.Printf("\n")

    }
}
	