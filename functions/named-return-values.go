package main
import "fmt"
func rectangle(len, wid float64) (area float64, perimeter float64) {
	area = len * wid
	perimeter = 2 * (len + wid)
	return
}
func main() {
	a, p := rectangle(5.5, 3.2)
	fmt.Println("THE AREA IS :", a)
	fmt.Println("THE PERIMETER IS:", p)
}