package main
import "fmt"
func sum(number ...int) int {
total := 0
for _, n := range number {
	total += n
}
return total
}
func main() {
	answer := sum(1, 2)
	fmt.Println(answer) //can also be written as fmt.Println(sum(1, 2)) directly instead of declaring the variable answer
}