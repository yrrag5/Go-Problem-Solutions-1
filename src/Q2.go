// Author: Garry Cummins
// Date: 01-10-2017

// Adapted from : https://tour.golang.org/welcome/4

package main
//Import time function to display time and date
import (
	"fmt"
	"time"
)

func main() {
    //Calling the time.Now function to display the current date and time 
	fmt.Println("The current date and time is: ", time.Now())
}