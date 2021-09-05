package homework

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

// GuessNumber 4.3 Đoán Số
func GuessNumber()  {
	number := randomNumber(0, 100)
	attempts := 0
	var guess int

	fmt.Println("Guess a number between 0 and 100.")

	for attempts < 10 {
		fmt.Println("Take a guess:")
		if _, err := fmt.Scanf("%d", &guess); err != nil {
			log.Fatalln("Error in Scan Attempts Number: ", err.Error())
		}
		attempts++
		if guess < number {
			fmt.Println("Too low.")
		}
		if guess > number {
			fmt.Println("Too high.")
		}
		if guess == number {
			break
		}
	}
	if guess == number {
		fmt.Printf("You guessed the number in %d tries\n", attempts)
	} else {
		fmt.Printf("The number was %d\n", number)
	}
}