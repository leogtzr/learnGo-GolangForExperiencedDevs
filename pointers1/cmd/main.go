package main

import "fmt"

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the parameter is: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}

func main() {
	var p *int32 = new(int32) //
	*p = 2

	fmt.Println(*p)

	var x int32 = 89
	p = &x
	*p += 1

	fmt.Println(*p)
	fmt.Println(x)

	var slice = []int32{7, 8, 9}
	var sliceCopy = slice
	sliceCopy[1] = 66
	fmt.Println(slice)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thing1 is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result)
}
