package main
import "fmt"
func main() {
	//1. Creation
	mySlice := []string{"GO", "STRING", "AWESOME"}
	fmt.Println("initial:", mySlice, "len:", len(mySlice), "cap:", cap(mySlice))
	//appending
	mySlice = append(mySlice, "programming")
	fmt.Println(mySlice, len(mySlice))
	dld := append(mySlice, "DG")
	fmt.Println(dld)
	//3. Slicing
	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	//4. modifying (note: this affect the underlying array)
	subSlice[0] = "REPLACED"
	fmt.Println(mySlice)
}