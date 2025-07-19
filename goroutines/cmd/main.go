package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait() // let's wait for the counter to be 0 ...
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 100_000_000; i++ {
		res++
	}
	wg.Done()
}
