package goroutine

import (
	"fmt"
	"time"
)

func Goroutine()  {
	
	ch := make(chan string)

	go func() {
		fmt.Println("Going to take a small nap..")
		time.Sleep(3*time.Second)
		ch <- "Hello"
	}()
	
	fmt.Println("Waiting for the message to come up .. ..")
	
	v:= <-ch

	fmt.Println("Recived Value form the channel",v)


}