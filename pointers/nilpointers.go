package main
import "fmt"
func main() {
	var ptr *int 
	// var num int = 42
	// ptr = &num//declared but not iniatilized - its nil
	fmt.Println("ptr is :", ptr) //<nil>
	// always check if ptr is nil before use
// *ptr = 10  // error:panic:runtime error: invalid memory address
if ptr != nil{
	fmt.Println(*ptr)
} else{
	fmt.Println("pointer is nil cannot use")
}
}