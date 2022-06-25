package faninandfanout

import (
	"fmt"
)

func FanOut(){
	ch1 , err := read("./FanIn_and_FanOut/file11.csv")
	if err != nil {
		panic(fmt.Errorf("could not read the file %v",err))
	}

	br1 := breakup("1",ch1)
	br2 := breakup("2",ch1)
	br3 := breakup("3",ch1)

	for{
		if br1 == nil && br2 == nil && br3 == nil{
			break
		}

		select {
			case _,ok := <-br1:
				if !ok{
					br1=nil
				}			
		
			case _,ok := <-br2:
				if !ok{
					br2=nil
				}
			case _,ok := <-br3:
				if !ok{
					br3=nil
				}				
		}		
	}
	fmt.Println("All completed cosuming and exiting")
}

func breakup(worker string,ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		for v:= range ch{
			fmt.Println(worker,"||",v,"||")
		}
		close(chE)
	}()
	return chE
}





