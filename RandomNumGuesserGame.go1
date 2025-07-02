package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randGenerator() {

	ran := rand.Intn(100)
	status := false
	fmt.Println("You only have 7 chances to get the correct number.")
	time.Sleep(4 * time.Second)
	fmt.Println("Rank is based on how much early you select the correct number")
	time.Sleep(4 * time.Second)

	fmt.Println("Game starts in:")
	time.Sleep(1 * time.Second)
	fmt.Println("3")

	time.Sleep(1 * time.Second)
	fmt.Println("2")

	time.Sleep(1 * time.Second)
	fmt.Println("1")

	time.Sleep(1 * time.Second)
	fmt.Println("GO")

	var guess int
	fmt.Scanln(&guess)

	for i := 0; i < 7; i++ {

		switch {

		case ran < guess:
			fmt.Println("lower")

		case ran > guess:
			fmt.Println("Higher")

		case ran == guess && i == 1 || i == 0:
			status = true
			time.Sleep(1 * time.Second)
			fmt.Println("Ohhh WHHAT?!")
			fmt.Println("Rank: The 'Verse'")
			fmt.Println("Tries:", i)

			if status == true {
				return
			}

		case ran == guess && i == 2:
			status = true
			time.Sleep(1 * time.Second)
			fmt.Println("The GUESS is tooo good")
			fmt.Println("Rank: Diamond")

			fmt.Println("Tries:", i)

			if status == true {
				return
			}

		case ran == guess && (i <= 3 || i <= 4):
			status = true
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Good.")
			fmt.Println("Rank: Gold")

			fmt.Println("Tries:", i)

			if status == true {
				return
			}

		case ran == guess && i == 5:
			status = true
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Ok.")
			fmt.Println("Rank: Platinium")

			fmt.Println("Tries:", i)

			if status == true {
				return
			}

		case ran == guess && i == 6:

			status = true
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Don't be to happy.")
			fmt.Println("Rank: Bronze")

			fmt.Println("Tries:", i)

			if status == true {
				return
			}

		}
		fmt.Scanln(&guess)

	}

	fmt.Println("The number was", ran)

}

func main() {
	randGenerator()
	fmt.Println("Play Again ? (y for YESS/anything else for NOOO)")
	var playAgain string
	fmt.Scanln(&playAgain)
	if playAgain == "y" {
		randGenerator()
	}
}
