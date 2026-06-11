package main
import "fmt"
func main() {
	var num int = 42
	var ptr *int
	ptr = &num
	fmt.Printf("type of ptr: %T\n", ptr)
	fmt.Println("value:", *ptr)
}