package goroutine

import (
	"fmt"
	"time"
)

 func BuffereChan()  {
	ch:= make(chan int, 2)

	go func() {
		for i:= 0 ; i < 3 ; i++ {
			fmt.Println(time.Now(),i,"*******sending*****")
			ch <- i
			fmt.Println(time.Now(),i,"----sent-----")
		}	
		fmt.Println(time.Now(),"!!!!!!!!All completed!!!!!!!!")
	}()
	time.Sleep(2*time.Second)
	fmt.Println("@@@@Waiting for the messages@@@@@@")
	fmt.Println(time.Now(),"^^^^^^received^^^^^^",<-ch)
	fmt.Println(time.Now(),"^^^^^^received^^^^^^",<-ch)
	fmt.Println(time.Now(),"^^^^^^received^^^^^^",<-ch)
	
	fmt.Println(time.Now(),")))))))))Exiting now from program((((((((((")
 }