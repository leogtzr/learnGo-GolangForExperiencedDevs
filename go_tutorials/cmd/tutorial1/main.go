package main

import "fmt"

func main() {
	var intNum int8
	intNum = 3
	fmt.Println(intNum)

	var myString string = "Hello World"
	fmt.Println(myString)

	var myName string = `Leo
	Gutierrez
	Ramirez`

	fmt.Println(myName)

	var myRune rune = 'a'
	fmt.Println(myRune)
}
