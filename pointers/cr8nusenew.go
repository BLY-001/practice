package main
import "fmt"
func main() {
	ptr := new(int)
	*ptr = 99
	fmt.Println("value :", *ptr)
}