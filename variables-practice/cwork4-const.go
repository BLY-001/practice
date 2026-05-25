package main
import "fmt"
func main() {
	const pi = 3.14159
	const AppName = "MyApp"
	const Maxuser = 100

	// alernatively multiple constants can be declared thus
	const(
		statusOK = 200
		StatusNotFound = 404
	)
	fmt.Println("pi =", pi)
	fmt.Println("my app name is :", AppName)
	fmt.Println("the maximum number of user is :", Maxuser)
	fmt.Println("the status is ok at :", statusOK)
	fmt.Println("the status is unacceptable at:", StatusNotFound)
}