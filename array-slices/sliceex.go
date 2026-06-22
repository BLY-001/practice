package main
import "fmt"
func main() {
	fruits := []string{"apple", "bannana", "cherries", "date", "elderberry"}
	slice1 := fruits[1:4]
	fmt.Println(slice1)
	slice2 := fruits[:3]
	fmt.Println(slice2)
	slice3 := fruits[2:]
	fmt.Println(slice3)
}