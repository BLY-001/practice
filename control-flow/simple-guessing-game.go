package main
import(
	"fmt"
	"math/rand"
	"time"
)
func main() {
	//seed random numbers
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(10) + 1 //random number from 1-10
	fmt.Printf("Guess a number between 1 and 10\n")
	for {
		var Guess int
		fmt.Println("your guess:")
		fmt.Scanln(&Guess)
		//check the guess with if-else
		if Guess < secret {
			fmt.Println("too low")
		} else if Guess > secret {
			fmt.Println("too high")
		} else {
			fmt.Println("correct!")
		}
		break //exit the loop
	}
}