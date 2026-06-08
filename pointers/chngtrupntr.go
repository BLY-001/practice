package main
import "fmt"
func main() {
	balance := 500
	fmt.Println("initial balance:", balance)
	ptr := &balance
	*ptr = 1000
	fmt.Println("new balance:", balance)
}