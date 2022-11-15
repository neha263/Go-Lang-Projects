package main

import (
	"fmt"
	"time"
)

func reader(ch chan string) {
	fmt.Println("value in ch : ", <-ch)
}
func main() {
	ch := make(chan string)
	go reader(ch)

	ch <- "Hello"
	close(ch)
	// ch <- "world" // Error : panic: send on closed channel

	time.Sleep(2 * time.Second)
}
