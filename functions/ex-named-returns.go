package main
import ("fmt"
"math")
func circle(radius float64) (diameter, circumference, area float64) {
diameter = 2 * radius
circumference = 2 * math.Pi * radius
area = math.Pi * radius * radius
return}
func main() {
	k, p, y := circle(4)
	fmt.Println(k)
	fmt.Println(p)
	fmt.Println(y)
	fmt.Printf("radius 5: diameter = %.2f, circumference = %.2f, area = %.2f\n", k, p, y)
}