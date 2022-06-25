package faninandfanout

import (
	"fmt"

	//"sync"
	"time"
)

func FanOut(){
	ch1 , err := read("./FanIn_and_FanOut/file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read the file %v",err))
	}

	ch2 , err := read("./FanIn_and_FanOut/file2.csv")
	if err != nil {
		panic(fmt.Errorf("could not read the file %v",err))
	}

	exit := make(chan struct{})


	chM := merge2(ch1,ch2)

	go func() {
	
		for v:= range chM{
			fmt.Println(time.Now(),"received","|||",v,"|||")
		}
		close(exit)		
	}()
	<-exit
}


// Merging into one channel using waitgroups like mutex concept 

// func merge1(cs ...<-chan []string) (<-chan []string) {
// 	var wg sync.WaitGroup
// 	out := make(chan []string)

// 	send := func(c <-chan []string)  {
// 		for n := range c{
// 			out <- n
// 		}
// 		wg.Done()
// 	}

// 	wg.Add(len(cs))

// 	for _ , c := range cs {
// 		go send(c)
// 	}

// 	go func(){
// 		wg.Wait()

// 		close(out)
// 	}()
// 	return out
// }


// Merging with chan buffer cocept
