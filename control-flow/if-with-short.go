package main
import "fmt"
func main () {
	if age := 20 ; age >= 18 {
		fmt.Println("adult")
	} //age is not avialble outside the if block
}