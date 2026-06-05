package main
import "fmt"
func safeDivide(a, b float64)(float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}else {
		return a/b, nil} 
}
func main() {
	value, err := safeDivide(5, 3)
	if err != nil {
		fmt.Print("Error:",err)
	} else {
		fmt.Println("the result is:", value)
	}
	value2, err2 := safeDivide(5, 0)
	if err2 != nil {
		fmt.Println("error:", err2)
	} else {
		fmt.Println("the second result is :", value2)
	}
}