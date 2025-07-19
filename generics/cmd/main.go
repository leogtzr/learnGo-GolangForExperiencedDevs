package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumSlice([]float32{2.0, 4.2, 6.2}))

	var values []bool = []bool{true, false}
	fmt.Println(isEmpty[bool](values))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
