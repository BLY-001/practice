package main
import "fmt"
func main() {
	var a, b float64
	var op string
	fmt.Print("the first number")
	fmt.Scanln(&a)
	fmt.Print("operator (+, -, /, *)")
	fmt.Scanln(&op)
	fmt.Print("the second number")
	fmt.Scanln(&b)
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0{
			fmt.Println("cannot divide by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("invalid operation")
		return
	}
	fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, result)
}