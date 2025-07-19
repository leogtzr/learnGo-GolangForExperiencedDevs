package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5)
	go process(c)
	for v := range c {
		fmt.Println(v)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
