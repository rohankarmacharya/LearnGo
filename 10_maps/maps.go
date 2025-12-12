package main

import (
	"fmt"
	"maps"
)

// map -> hash, object, dictionary
func main() {

	//create map
	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["area"] = "backend"

	// //get value
	// fmt.Println(m["name"], m["area"])

	// Note: If key doesn't exist in the map, it returns zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 100
	// fmt.Println(m["phone"]) // prints 0

	// // fmt.Println(len(m)) //	prints 2
	// // delete(m, "price")
	// // fmt.Println(m)

	// clear(m) // clears the value of map
	// fmt.Println(m)

	// m := map[string]int{"age": 30, "price": 100}

	// v, ok := m["age"] // check if key exists
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("key exists")
	// } else {
	// 	fmt.Println("key doesn't exist")
	// }

	m1 := map[string]int{"age": 30, "price": 100}
	m2 := map[string]int{"age": 30, "price": 100}

	fmt.Println(maps.Equal(m1, m2))
}
