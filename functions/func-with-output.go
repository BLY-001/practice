package main
import "fmt"
func add(x int, y int) int {
	result := x + y
	return result
}
func main() {
	answer := add(5, 3)
	fmt.Println(answer)
}