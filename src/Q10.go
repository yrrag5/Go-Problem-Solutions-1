// Author: Garry Cummins
// Date: 06/10/2017

// Adapted from: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main 
import "fmt"

func main() { 

	
        input := "knalB"
        n := 0

		//Takes in user string input 
        fmt.Println("Please input a word...")
        fmt.Scanf("%s", &input)
        
		//Using rune to transform a word into unicode
        rune := make([]rune, len(input))
        for _, i := range input { 
                rune[n] = i
                n++
        } 
        rune = rune[0:n]
        // Rewrites the numbers for it to be read backwards
        for j := 0; j < n/2; j++ { 
                rune[j], rune[n-1-j] = rune[n-1-j], rune[j] 
        } 
        // Reveres the word back into intger values 
		//Will only reverse the first word if more than one is entered
        output := string(rune)
        fmt.Println(output)
}