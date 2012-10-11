package main

import (
	"fmt"
	"time"
	"math/rand"
	)

func main(){
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s:= <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're talk too much.")
			return
		}
		
	}
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

func boring(msg string) <-chan string{ // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) 
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}