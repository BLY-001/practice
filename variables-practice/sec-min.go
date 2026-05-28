package main
import "fmt"
func main() {
	seconds := 3665
	minutes := seconds/60
	s := seconds % 60
	fmt.Printf("%dsecs = %dmins,%dsecs\n", seconds, minutes, s)
}