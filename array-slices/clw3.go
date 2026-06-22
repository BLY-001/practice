package main
import "fmt"
func main() {
	s := make([]int, 3, 5)
	fmt.Printf("initial length =%d and the capacity = %d\n", len(s), cap(s))
	for i := 0; i < 5; i++ {
		s = append(s, i)
fmt.Printf("after append%d final length=%d and the final capacity= %d\n", i, len(s), cap(s))	
}
}