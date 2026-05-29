package main
import "fmt"
func main() {
	score := 65
if score >= 70 {
	fmt.Println("excellent : Grade A")
} else if score >= 60 {
	fmt.Println("very good: Grade B")
} else if score >= 50 {
	fmt.Println("good : Grade C")
} else {
	fmt.Println("failed, advice to retake")
}
}