package main

import (
	"fmt"
	"time"
)

func read_value(ch chan int) {
	fmt.Println("value in channel :", <-ch)
}

func main() {
	ch := make(chan int)
	go read_value(ch)

	ch <- 123
	time.Sleep(3 * time.Second)
}
