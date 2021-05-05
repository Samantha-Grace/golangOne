package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix()) //Initializing the random generator.

	var (
		c       int
		gameVal = rand.Intn(100) // 0 <= gameVal < 100
	)

	for {
		fmt.Println("enter a number(enter -1 to exit):")
		fmt.Scanf("%d", &c)
		if c == -1 {
			break // breaks the infinite loop and goes to the first line after loop finishing bracket.
		}

		if c < gameVal {
			fmt.Println("your guess is little than the game value ")
		} else if c > gameVal {
			fmt.Println("your guess is greater than the game value ")
		} else if c == gameVal {
			fmt.Println("Bingo! you got it ")
			fmt.Println("random number was:", gameVal)
			fmt.Println("your input was:", c)
			break
		}
	} // end of infinite loop.

}
