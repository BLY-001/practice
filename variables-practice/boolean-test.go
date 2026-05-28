package main
import "fmt"
func main() {
	age := 20
	hasLicense := true
	canDrive := hasLicense && age >= 18
	fmt.Println("can he drive ?", canDrive)
}