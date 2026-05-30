package main
import "fmt"
func main() {
	// traditional loop (start ; condition ; step)
	for i := 0 ; i < 5 ; i++ {
		fmt.Println(i)
	}
	// while-like loop(condition only)
	count := 0
	for count < 5 {
		fmt.Println("total count is :", count)
		count++
	}
	// infinite loop(with break to exit)
	for {
		fmt.Println("loop forever unless we break")
		break
	}
}