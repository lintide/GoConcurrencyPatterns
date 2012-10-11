package main

import (
	"fmt"
	// "time"
	"math/rand"
	)

func main(){
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(20); i >= 0; i-- { fmt.Println(<-c) }
	quit <- true
	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
				case s := <-input1: c <- s
				case s := <-input2: c <- s
			}
		}
	}()
	return c
}

func boring(msg string, quit chan bool) <-chan string{ // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
			
			//time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}