package main
import "fmt"
func main() {
	var n int
	fmt.Println("the value of N:")
	fmt.Scanln(&n)
	sum := 0
	for i := 1 ; i <= n ; i++ {
		sum += i
	}
	fmt.Printf("sum 1 to %d = %d\n", n, sum)
}