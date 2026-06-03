package main
import "fmt"
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "cannot divide by zero"
	}
	return a/b, ""
}
func main() {
	result, errMsg := divide(10, 2)
	if errMsg != "" {
		fmt.Println("Error:", errMsg)
	} else {
		fmt.Println("result:", result)
	}

	result2, errMsg2 := divide(10, 0)
	if errMsg2 != "" {
		fmt.Println("Error:", errMsg2)
	} else {
		fmt.Println("result:", result2)
	}
}