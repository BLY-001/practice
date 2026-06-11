package main
import "fmt"
func main() {
	balance := 500
	fmt.Println("initial balance:", balance)
	ptr := &balance //ptr points to balance
	*ptr = 1000 // change the value at that address to 1000
	fmt.Println("new balance:", balance)
}