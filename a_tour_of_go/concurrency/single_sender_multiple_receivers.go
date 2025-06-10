package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int, 10)

	// single sender
	go func(){
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// multiple receivers
	for i := 0; i < 3; i++ {
		// each iteration starts a new receiver goroutine
		// the receivers read from the channel in parallel
		// so the tasks are randomly distributed
		go func(id int){
			for v := range ch {
				fmt.Printf("Receiver %d received task %d\n", id, v)
			}
		}(i)
	}

	// goroutines run asynchronously 
	// the program exits right after starting the goroutines
	// sleep gives the goroutines time to print their output 
	// before the program exits
	time.Sleep(1 * time.Second)
}