package main

import "fmt"

func main() {
	// if else statement

	// age := 18

	// if age >= 18 { //round braces are not required
	// 	fmt.Println("A person is an adult")
	// } else {
	// 	fmt.Println("A person is under age.")
	// }

	// age := 16
	// if age >= 18 {
	// 	fmt.Println("A person is an adult.")
	// } else if age >= 12 {
	// 	fmt.Println("A person is a teenager.")
	// } else {
	// 	fmt.Println("A person is a child.")
	// }

	//use of logical operators

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }

	if age := 15; age >= 18 {
		fmt.Println("A person is an adult.", age)
	} else if age >= 12 {
		fmt.Println("A person is a teenager.", age)
	} else {
		fmt.Println("A person is a child.", age)
	}

	//go does not have ternary operator
}
