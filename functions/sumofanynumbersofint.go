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
	fmt.Println(sum(1, 2)) //can also be written as answer := sum(1, 2) and then fmt.Println(answer)
fmt.Println(sum(5, 10, 15))
fmt.Println(sum())
}