package main

//Buffered Channels allow sending multiple values without blocking until full

import "fmt"

func main(){
	ch:= make(chan int, 2)
	ch<-1
	ch<-2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//Buffered Channels dont block immediately unless full.
//Reading from an empty channel blocks.
//If the buffer is full, further sends block until space is available.
