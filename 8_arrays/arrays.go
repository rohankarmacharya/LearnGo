package main

import "fmt"

func main() {
	//zeroed values
	// int -> 0, bool -> false, string -> ""

	// var arr [4]int

	// arr[0] = 10

	// fmt.Println(arr[0])  // prints 10
	// fmt.Println(arr) //prints [10 0 0 0]
	// fmt.Println(len(arr))   // prints length of array

	// var vals [4]bool // prints [false false false false]
	// vals[2] = true   // assigns true to index 2
	// fmt.Println(vals)

	// var names [3]string
	// names[0] = "golang"

	// fmt.Println(names)

	//array literal
	// numbers := [4]int{1, 2, 3, 4}
	// fmt.Println(numbers)

	//2D array
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)

}
