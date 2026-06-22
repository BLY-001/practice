package main
import "fmt"
func main() {
	a := []int{1,2,3}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 99
	fmt.Println(a)//[1 2 3] unchanged
}