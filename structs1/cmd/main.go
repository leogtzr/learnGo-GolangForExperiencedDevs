package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
}

type gasEngine2 struct {
	mpg     uint8
	gallons uint8
	owner
}

func main() {
	var myEngine gasEngine
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{mpg: 2, gallons: 3}
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine = gasEngine{2, 3, owner{"Leo"}}
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine3.mpg, myEngine3.gallons)

	myEngine3.gallons = 78

	ge := gasEngine2{mpg: 4, gallons: 9, owner: owner{name: "Leo"}}
	fmt.Printf("%+v\n", ge)
	fmt.Println(ge.name)
}
