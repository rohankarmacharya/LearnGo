package main

import (
	"fmt"
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
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")

	// default:
	// 	fmt.Println("It's work day")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I'm an integer and my value is %d\n", v)
		case string:
			fmt.Printf("I'm a string and my value is %s\n", v)
		default:
			fmt.Println("I'm an unknown type", v)
		}
	}
	whoAmI(7)
	whoAmI("golang")
	whoAmI(true)
}
