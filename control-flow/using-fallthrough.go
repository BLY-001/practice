package main
import "fmt"
func main() {
	price := 2
	switch price{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough // fallthrough only works for the case immediately after it and the rest of the code runs normallly
	case 3:
		fmt.Println("nthree")
	default:
		fmt.Println("three")
	}
}