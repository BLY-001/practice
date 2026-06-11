package main
import "fmt"
func double(n int) {
	n = n * 2
	fmt.Println("inside function", n)
}
func main() {
	value := 10
	double(value)
	fmt.Println("outside functions:", value)
}