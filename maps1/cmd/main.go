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

	ages["maria"] = 55

	for name, age := range ages {
		fmt.Printf("%s has %d age\n", name, age)
	}

	var nums [3]int = [3]int{33, 34, 29}
	for i, v := range nums {
		fmt.Printf("i(%d) = %d\n", i, v)
	}

	// -------
	var i int = 0
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}

	fmt.Println("---------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
