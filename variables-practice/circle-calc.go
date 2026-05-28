package main
import (
	"fmt"
	"math"
)
func main() {
r := 7.5
diameter := (2 * r)
circumference := (math.Pi * diameter)
area := (math.Pi * r * r)
fmt.Printf("the diamter of this circle is %.1f\n", diameter)
fmt.Printf("the circumference of this circle is %.1f\n", circumference)
fmt.Printf("the area of this circle is %.1f\n", area)
}