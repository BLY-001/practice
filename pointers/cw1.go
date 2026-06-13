package main
import "fmt"
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	x, y := 10, 20
	fmt.Println("before swap:", x, y)
	swap(&x, &y)
	fmt.Println("after swap", x, y)
}