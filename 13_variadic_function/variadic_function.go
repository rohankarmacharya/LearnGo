package main

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	nums := []int{1, 2, 3}
	result := sum(nums...)
	println(result)
}
