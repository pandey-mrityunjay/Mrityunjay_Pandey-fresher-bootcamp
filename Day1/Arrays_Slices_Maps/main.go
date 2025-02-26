package main

import "fmt"

// array is a fixed size sequence of elements
func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	slices_usage()
	map_usage()
	slices_arr()
}

// Slice is a dynamically-sized sequence of elements
func slices_usage() {
	nums := []int{10, 20, 30}
	nums = append(nums, 40, 50)

	fmt.Println(nums)
	//unlike arrays , slices grow dynamically with append
}

// Maps: is a collection of key-value pairs
func map_usage() {
	studentMarks := make(map[string]int)

	studentMarks["Alice"] = 80
	studentMarks["Bob"] = 90

	fmt.Println(studentMarks)
}

// iterating over a slice
func slices_arr() {
	number := []int{10, 20, 30}
	for index, value := range number {
		fmt.Println(index, value)
	}
}
