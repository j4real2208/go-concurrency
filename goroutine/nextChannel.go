package goroutine

import (
	"fmt"
	"time"
)

func NextChan(){
	ch:= make(chan int , 2 )
	exit:= make(chan struct{})

	go func() {
		for i:= 0 ; i <= 5 ; i++ {
			fmt.Println(time.Now(),"|||",i,"|||","*******sending*****")
			ch <- i
			fmt.Println(time.Now(),"|||",i,"|||","----sent-----")

			time.Sleep(1*time.Second)
		}
		fmt.Println(time.Now(),"!!!!!!!!All completed!!!!!!!!")
		close(ch)
	}()
		// Select method in chan
	// go func() {
	// 	for{
	// 	 select {
	// 	 	case v, open := <- ch:
	// 			if !open {
	// 				close(exit)
	// 				return
	// 			}
	// 			fmt.Println(time.Now(),"received","|||",v,"|||")
	// 		//default:
	// 			//fmt.Println("NO messages in channel to recive")
	// 	 }
	// 	}
	// }()
	for v:= range ch {
		fmt.Println(time.Now(),"received","|||",v,"|||")
	}
	close(exit)
	fmt.Println(time.Now(),"Waiting for everything to complete")
	<-exit
	fmt.Println(time.Now(),"Completed the transaction")
}