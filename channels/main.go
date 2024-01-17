package main

import (
	"fmt"
	"time"
)

func main() {

	msgchn := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		msg := <-msgchn
		fmt.Println(msg)

	}()

	msgchn <- 20

}
