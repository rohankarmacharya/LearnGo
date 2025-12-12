package main

import "fmt"

func counter() func(int) int {

	var count int = 0
	return func(x int) int {
		count += x
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment(1))
	fmt.Println(increment(1))
}
