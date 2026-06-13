package main
import "fmt"
func main() {
	x := new(int)
	*x = 99
	fmt.Println(*x)
}