package main
import "fmt"
func main() {
	//the use of break stops the printing at a particular point i.e it stops at 4  
	for i := 0 ; i <= 10 ; i++ {
		if i == 5{
			break
		}
		fmt.Println(i)
	}
	// the use of continue skips the particular point and still continue i.e it skips j = 5 and prints the rest
	for j := 0 ; j <= 10 ; j++ {
		if j == 5{
			continue
		}
		fmt.Println(j)
	} 
}