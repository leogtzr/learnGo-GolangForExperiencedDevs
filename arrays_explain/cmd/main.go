package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{5, 25, 45}
	fmt.Printf("The values are: %+v\n", intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	intOther := [3]int32{5, 6, 2}
	fmt.Println(intOther)

	intOther2 := [...]int32{5, 2, 89}
	fmt.Println(intOther2)
}
