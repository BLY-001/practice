package main
import "fmt"
func main() {
	original := []int{1,2,3,4,5}
	sub := original[1:4]
	fmt.Printf("the initial original is %d\n", original)
	sub[1] = 99
	fmt.Println("the new original is :", original)

}