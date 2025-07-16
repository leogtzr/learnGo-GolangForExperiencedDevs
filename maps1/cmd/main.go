package main

import "fmt"

func main() {
	dict1 := make(map[string]int)
	fmt.Println(dict1)

	dict1["leonardo"] = 34
	fmt.Println(dict1)

	fmt.Println(dict1["leonardo"])

	ages := map[string]uint8{"Perla": 33, "Leonardo": 34}
	fmt.Println(ages)

	if mariaAge, ok := ages["maria"]; ok {
		fmt.Println(mariaAge)
	}

}
