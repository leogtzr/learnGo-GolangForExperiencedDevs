package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed) // 114, uint8

	// index +
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Println(myRune) // 97

	var strBuilder strings.Builder
	strings := []string{"Leonardo", "Gutierrez", "Ramirez"}

	for i := range strings {
		strBuilder.WriteString(strings[i])
	}

	fmt.Printf("The string is: '%s'\n", strBuilder.String())
}
