package main
import "fmt"
func main() {
	//variable declaration
	var a, b int
	a = 15
	b = 4
	//operations to perform
	sum := a + b
	difference := a - b
	multiply := a * b
	quotient := a / b //integer division(no decimal)
	remainder := a % b //modulo (remainder)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d * %d = %d\n", a, b, multiply)
	fmt.Printf("%d / %d = %d\n", a, b, quotient)
	fmt.Printf("%d - %d = %d\n", a, b, difference)
	fmt.Printf("%d %% %d = %d\n", a, b, remainder)
}