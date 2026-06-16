package main
import "fmt"
//declaring an array of 5
var scores [5]int
//the array is fixed , you cannot make it hold 6 later

//setting and getting  values 
func main() {
	var days [7] string
	days[0] = "sunday"
	days[1] = "monday"
	days[2] = "tuesday"
	days[3] = "wednesday"
	days[4] = "thursday"
	days[5] = "friday"
	days[6] = "saturday"
	fmt.Println(days[0])
	fmt.Println(days[5])
	//shorter initialization 
	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes[2])
	// you can let Go count the length for you:
	colors := [...]string{"red", "green", "blue"}
	fmt.Println(colors[2])
}