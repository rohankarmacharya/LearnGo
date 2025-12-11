package main

import "fmt"

// var name string = "golang" // allowed outside function scope
// name := "golang" // not allowed outside function scope

func main() {
	// var name string = "golang"

	var name = "golang" // golang infers type string

	fmt.Println(name)

	var is_adult = true // golang infers type bool
	fmt.Println(is_adult)

	var age int = 9
	fmt.Println(age)

	//shorthand syntax
	lang := "golang"
	fmt.Println(lang)

	// value pachi assign garne ho bhane type define garnai parcha
	var score int
	score = 100
	fmt.Println(score)

	// var price float32 = 99.99
	// var price = 99.99 // golang infers type float64
	price := 99.99
	fmt.Println(price)
}
