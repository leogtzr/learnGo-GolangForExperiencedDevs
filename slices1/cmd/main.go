package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{4, 3, 2, 1}
	fmt.Println(intSlice)

	intSlice = append(intSlice, 0)
	fmt.Println(intSlice)

	// Length and capacity
	fmt.Printf("Length: %d, capacity: %d\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)

	fmt.Println(intSlice)

	var intSlices3 []int32 = make([]int32, 5)
	intSlices3[0] = 4
	fmt.Println(intSlices3)
}
