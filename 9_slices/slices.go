// slices are dynamic arrays
// most used construct in go
// uninitialized slice is nil

package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int
	// fmt.Println(numbers == nil) // prints true
	// fmt.Println(len(numbers))   // prints 0

	// var numbers = make([]int, 0, 5) // length 3, capacity 5
	// // fmt.Println(numbers)        // prints [0 0 0]
	// // fmt.Println(numbers == nil) // prints false
	// // fmt.Println(len(numbers))   // prints 3
	// // fmt.Println(cap(numbers))         // Capacity : prints 3  //Capacity is total space allocated; dynamically resizes
	// numbers = append(numbers, 10)     // adds 10 to the end of slice
	// numbers = append(numbers, 11, 12) //capacity increases automatically
	// fmt.Println(numbers)              // prints [0 0 0 10]
	// fmt.Println(cap(numbers))
	// fmt.Println(len(numbers))

	// var numbers = make([]int, 0, 5)
	// numbers = append(numbers, 2)

	// var numbers2 = make([]int, len(numbers))
	// copy(numbers2, numbers) // copy function

	// fmt.Println(numbers, numbers2)

	// slice operator
	// numbers := []int{1, 2, 3}
	// fmt.Println(numbers[0:2]) // prints [1 2]
	// fmt.Println(numbers[1:]) // prints [2 3]

	//slices function
	var numbers1 = []int{1, 2, 3, 4}
	var numbers2 = []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(numbers1, numbers2))

	//2D slice
	numbers3 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(numbers3)

}
