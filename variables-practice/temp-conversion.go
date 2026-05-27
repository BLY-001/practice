package main
import "fmt"
func main() {
	var Celsius float64
	Celsius = 23
	Farhrenheit := (1.8 * Celsius) + 32
	fmt.Printf("%.2f°C = %.2f°F\n", Celsius, Farhrenheit)
	var F float64 = 80
	C := (F - 32) / 1.8
	fmt.Printf("%.2f°F = %.2f°C\n", F, C)
}