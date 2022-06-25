package main

import (
	"fmt"

	//"github.com/j4real2208/go-concurrency/goroutine"
	fn "github.com/j4real2208/go-concurrency/FanIn_and_FanOut"
)

func main() {
	//go PrintHell()
	//fmt.Println("Here you go print ")	

	//goroutine.Goroutine()
	
	// -> Buffered Channel and read 
	//goroutine.BuffereChan()
	
	//-> Buffered channel and synchronized 
	//goroutine.NextChan()

	// -> Fan In Method read two files and output through single channel
	fn.FanOut()


}


func PrintHell(){
	fmt.Println("HEllo world this is not going to be print")
}