package main
import "fmt"
func main() {
 var ask int
 fmt.Println("put a number here:")
 fmt.Scanln(&ask)
 if ask < 0 {
	fmt.Println("NEGATIVE")
 } else if ask > 0 {
	fmt.Println("POSITIVE")
 } else {
	fmt.Println("number is zero")
 }
}