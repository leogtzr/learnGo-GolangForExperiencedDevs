package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type owner struct {
	name string
}

type gasEngine2 struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

func main() {
	var myEngine gasEngine
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{mpg: 2, gallons: 3}
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine = gasEngine{2, 3, owner{"Leo"}}
	fmt.Printf("mgp: %d, gallons: %d\n", myEngine3.mpg, myEngine3.gallons)

	myEngine3.gallons = 78

	ge := gasEngine2{4, 9, owner{"Leo"}, 3}
	fmt.Printf("%+v\n", ge)
	fmt.Println(ge.name)
	fmt.Println(ge.int)

	owner := struct {
		name string
		age  int
	}{name: "Leonardo", age: 34}

	fmt.Println(struct{ name string }{name: "Leonardo"}.name)

	fmt.Println(owner)

	fmt.Println(myEngine3.milesLeft())
}
