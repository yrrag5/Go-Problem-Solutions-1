// Author: Garry Cummins
// Date: 06/10/2017

// Adapted from: https://gist.github.com/abesto/3476594

package main

import (
	"fmt"
	"math"
)

//Newtons method to discover the square root 
func z_next(z float64, x float64) float64{
	return z - (((z * z) - x) / (2 * z))
}

func main() {
	x := 20.0
	
	var z float64
	
	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %12.2f\n", z)
	}

	// Square root based on newtons method 
	fmt.Printf("The square root of %.2f is %.2f \n", x, z)
	// Math import shows exact root 
	fmt.Printf("Math import calculation: %.2f", math.Sqrt(x));
}
