//when you have a pointer to astruct,go allows you to use . directly (no need to write (*ptr).field - go does it aautomtically)

package main
import "fmt"
type Person struct{
	name string
	age int
}
func main() {
	p := Person{name: "alice", age: 30}
	ptr := &p
	(*ptr).name = "bob" //without shortcut also works but ugly
	ptr.age = 30
	fmt.Println(p)
}