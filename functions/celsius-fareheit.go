package main
import "fmt"
func toFahrenheit(celsius float64) float64{
	return (celsius * 1.8) + 32
}
func main() {
	temps := [] float64 {0, 25, 100}
	for _, c := range temps {
		fmt.Printf("%.1f°C = %.1f°F\n", c, toFahrenheit(c))
	}
	// fmt.Println(toFahrenheit(0))
	// fmt.Println(toFahrenheit(25))
	// fmt.Println(toFahrenheit(100))
}