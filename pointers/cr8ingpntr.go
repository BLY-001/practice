package main
import "fmt"
func main() {
	score := 100
	ptr := &score
	fmt.Println("score value:", score)
	fmt.Println("ptr (address):", ptr)
	fmt.Println("*ptr (value at that address):", *ptr)
}