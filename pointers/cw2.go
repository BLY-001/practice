package main
import "fmt"
func double(n *int) {
	*n = *n * 2
}
func main() {
	value := 7
	fmt.Println(value)
	double(&value)
	fmt.Println(value)
}
