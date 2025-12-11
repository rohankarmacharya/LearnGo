package main

import "fmt"

// for loop is only way to do looping in golang
func main() {

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinite loop
	// for {
	// 	println("1")
	// }

	//classic for loop
	for i := 0; i <= 3; i++ {
		// break // to exit loop

		if i == 2 {
			continue // to skip current iteration
		}
		fmt.Println(i)
	}

	//range
	for i := range 3 { //prints 0 1 2
		fmt.Println(i)
	}

}
