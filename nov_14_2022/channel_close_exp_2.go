package main

import (
	"fmt"
	"time"
)

func reader(ch chan int) {
	for {
		val, ok := <-ch
		if ok == false {
			fmt.Println("loop break")
			break
		} else {
			fmt.Println(val, ok)
		}
	}
}
func main() {
	ch := make(chan int)

	go reader(ch)

	ch <- 10
	close(ch)
	ch <- 20

	time.Sleep(time.Second * 3)
}
