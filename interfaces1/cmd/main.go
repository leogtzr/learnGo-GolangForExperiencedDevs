package main

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	return miles <= e.milesLeft()
}

func main() {
}
