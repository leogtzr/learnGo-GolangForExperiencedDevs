package main

import (
	"errors"
	"fmt"
)

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

	// How to define a variable
	name1 := "Leonardo"
	var name2 = "Leonardo"
	var name3 string = "Leonardo"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	var var1, var2 = 1, 2
	var3, var4 := 1, 2

	fmt.Println(var1, var2, var3, var4)

	someVal, _ := myFunc()
	fmt.Println(someVal)

	res, remainder, err := intDivision(10, 3)
	if err != nil {
		fmt.Printf("There was an error: ")
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", res)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", res, remainder)
	}

	// Now using the switch clause
	switch {
	case err != nil:
		fmt.Printf("There was an error: ")
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", res)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", res, remainder)
	}

	fmt.Printf("The result is: %d\n", res)
}

func myFunc() (int, string) {
	return -1, "Leo"
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("divide by zero")
		return -1, -1, err
	}

	var result int = numerator / denominator
	remainder := numerator % denominator

	return result, remainder, nil
}
