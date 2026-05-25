package main
import "fmt"
func main() {
	var x, y float64
	x = 15.0
	y = 4.0

	addition := x + y
	division := x / y
	

	fmt.Printf("%.3f + %.3f = %.3f\n", x, y, addition)
	fmt.Printf("%.3f / %.3f = %.3f\n", x, y, division)

}