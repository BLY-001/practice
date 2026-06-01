package main
import "fmt"
func main() {
	fmt.Println("imput a number for the log")
	var number int
	fmt.Scanln(&number)
	for i := 1 ; i <= 10 ; i++ {
		var result int = number * i
		fmt.Printf("%d * %d = %d\n", number, i, result)
	}
}