package main

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string, bool) {
// 	return "golang", "python", "java", true
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	// result := add(3, 6)
	// fmt.Println("Sum is:", result)
	// lang1, lang2, lang3, bool1 := getLanguages()
	// fmt.Println(lang1, lang2, lang3, bool1)

	fn := processIt()
	result := fn(5)
	println(result)

}
