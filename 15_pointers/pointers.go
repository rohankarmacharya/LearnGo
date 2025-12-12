package main

import "fmt"

func changeNum(x *int) {
	*x = 10
	fmt.Println("In changeNum:", *x)
}

func main() {
	x := 1
	changeNum(&x)
	fmt.Println("After changeNum:", x)
}
