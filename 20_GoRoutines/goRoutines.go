package main

// go routines are lightweight threads in Go
// They allow concurrent execution of functions

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}
func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}
	// Add sleep to allow goroutines to complete
	time.Sleep(time.Second * 2)
	fmt.Println("All tasks completed")
}
