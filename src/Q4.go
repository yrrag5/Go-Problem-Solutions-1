// Author: Garry Cummins
// Date: 01-10-2017

// Adapted from : https://play.golang.org/p/5lZMq3bd8a

package main

import (
	"fmt"
	"math/big"
)
//Calculation for finding the factorial of a number
func factorial(n int64) *big.Int {
   if n < 0 {
      return big.NewInt(1)
   }
   if (n==0) {
      return big.NewInt(1)
   }
   bigN := big.NewInt(n)
   return bigN.Mul(bigN, factorial(n-1))
}

func main() {
	//Set number value for factorial 
   fmt.Println(factorial(100))
}