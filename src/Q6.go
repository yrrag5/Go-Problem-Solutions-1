// Author: Garry Cummins
// Date: 06/10/2017

// Adapted from: https://gist.github.com/pyk/10519339

package main
import (	
	"fmt"
)

func main() {
	// Variables with array of random integer values 
	var temp,s, l int

	array := []int{
		8,16,24,32,
		5,10,15,20,
		7,14,21,28,
		100,1,88,27,
	}

	// Gets every value in the hard coded array 
	fmt.Printf("This is the array of numbers: %v \n\n", array)

	//Compares every integer in the array until the largest value is found 
  	for _,c:=range array {
		if c > temp {
    		fmt.Println(c,">",temp)
    		temp = c
    		l = temp
    	} else {
      	fmt.Println(c,"<",temp)
    	}
  	}
	//Prints largest number   
	fmt.Println("The largest number in the array is: ", l)

	//Compares every integer value in the array until the smallest value is found
	for _,c:=range array {
		// Starting value for c will be the highest number recently found 
		if c > temp {
			fmt.Println(c,">",temp)
		} else {
			fmt.Println(c,"<",temp)
			temp = c
			s = temp
		}
	}
	//Prints smallest number 
	fmt.Println("The smallest number in the array is: ", s)
}