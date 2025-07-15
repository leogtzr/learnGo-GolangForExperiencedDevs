package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{4, 3, 2, 1}
	fmt.Println(intSlice)

	intSlice = append(intSlice, 0)
	fmt.Println(intSlice)

	// Length and capacity
	fmt.Printf("Length: %d, capacity: %d\n", len(intSlice), cap(intSlice))
}
