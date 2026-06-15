package main
import "fmt"
type student struct{
	name string
	age int
}
func havebirthday(s *student) {
	s.age++
}
func main() {
stu := student{name: "binlawal", age: 20}
fmt.Println("before birthday:", stu)
havebirthday(&stu)
fmt.Println("after birthday", stu)
}