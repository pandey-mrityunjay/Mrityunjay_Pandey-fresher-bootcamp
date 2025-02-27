package main

import (
	"fmt"
	"time"
)

func sender1(ch chan string){
	time.Sleep(2*time.Second)
	ch<-"Message from sender1"
}

func sender2(ch chan string){
	time.Sleep(1*time.Second)
	ch<-"Message from sender2"
}

func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)

	go sender1(ch1)
	go sender2(ch2)

	select{
	case msg :=<-ch1:
		fmt.Println("Received:",msg)
	case msg := <-ch2:
		fmt.Println("Received:",msg)
	}
}