package main
import "fmt"
func add(x int, y int) { //can also be written as func add(x, y int) {
sum := x + y
fmt.Println(x, "+", y, "=", sum)
}
func main() {
	add(5, 4)
	add(5, 1,)
}