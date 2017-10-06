// Author: Garry Cummins
// Date: 01-11-2017

// Adapted from : http://golang-basic.blogspot.ie/2014/05/guess-number-game-in-go.html

package main

 import (
         "fmt"
         "math/rand"
         "time"
 )

 func numRand(min, max int) int {
         rand.Seed(time.Now().UTC().UnixNano())
         return rand.Intn(max-min) + min
 }

 func main() {
	 	// Variables
	    var guess int
		guessTaken := 0
		i := 1
		prevGuess := 0	
        num := numRand(1, 100)

		// Will loop until the user guesses the correct number 
		// Accepts string values due to no check 
		for{

			fmt.Println("Guess a number between 1 and 100: ")
			fmt.Scanf("%d", &guess)

			// Flushes the buffer to prevent repeats of question 
			fmt.Scanf("%d")
			guessTaken++

			// Checks if the user inputed a repeated number so it doesnt get counted 
			if guess == prevGuess{
				fmt.Println("The number was already used")
				i--
			}
			// Gives user estimates of how close they are to guessing the number 
			if guess < num {
				fmt.Println("Your guess is too low!")
			}

			if guess > num {
				fmt.Println("Your guess is too high!")
			}

			if guess == num {
				// Prints number and number of attempts it took the user to guess it 
				fmt.Println("\nNumber is :", num)
				fmt.Printf("Correct! You guessed the number in %d tries\n", guessTaken)
				return 
			}
			i++
			prevGuess = guess
			
		}//For
 }
