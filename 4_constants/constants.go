package main

import "fmt"

const name = "golang" // allowed outside function scope
// name := "golang" // not allowed outside function scope

func main() {
	// const name string = "golang"
	// const name = "golang" // type inferred as string
	// name = "python" // error: cannot assign to name (constant variable)
	fmt.Println(name)

	// constants grouping
	const (
		port = 5000
		host = "localhost"
	)

	// port = 6000 // error: cannot assign to port (constant variable)
	fmt.Println(port, host)
}
