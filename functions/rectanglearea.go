package main
import "fmt"
func area(length, width float64) float64 {
	area := length * width
	return area
}
func main() {
	par := area(5.5, 3.2)
	fmt.Println(par)
}