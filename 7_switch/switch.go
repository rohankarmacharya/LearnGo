package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 4

	// //simple switch
	// switch i {
	// case 1:
	// 	fmt.Println("one")

	// case 2:
	// 	fmt.Println("two")

	// case 3:
	// 	fmt.Println("three")

	// default:
	// 	fmt.Println("unknown")
	// }

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")

	default:
		fmt.Println("It's work day")
	}

}
