//iterating over a slice with for range
package main
import "fmt"
func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names{
		fmt.Printf("index %d: %s\n", index, name)
	}
	//if you only need the value, use underscore for the index i.e
	domes := []string{ "law", "stud", "lan"}
	for _, dome := range domes{
		fmt.Println(dome)
	}
	//if you want to check if slice is empty
	var items []string
	if len(items) == 0 {
		fmt.Println("slice is empty")
	}
	//removing an element(by index)
	nums := []int{20,30,40,50,60}
	i := 2
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums)
	//inserting an element(by index)
	nums = append(nums[:i], append([]int{40}, nums[i:]...)...)
	fmt.Println(nums)
}