package main
import "fmt"
func increment(c *int) {
*c = *c + 1
}
func main() {
	count := 0
	for i := 0; i < 5; i++{
		increment(&count)
	}
	fmt.Println(count)
}