package main

import "fmt"

func main() {
	var p *int32 = new(int32) //
	*p = 2

	fmt.Println(*p)
}
