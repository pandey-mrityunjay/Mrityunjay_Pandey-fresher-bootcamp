package main

import (
	"fmt"
	"time"
)

func sender(id int, ch chan string) {
	//Each sender will send 3 messages with a delay proportional to their ids
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("Sender %d: Message %d", id, i)
		time.Sleep((time.Duration(id) * time.Second))

	}
	close(ch)//Important to close the channel to stop the loop
}

func main(){
	//Two unbuffered channels are created
	//Will be used by sender goroutines to send messages
	ch1:=make(chan string)
	ch2:=make(chan string)

	go sender(1,ch1)
	go sender(2,ch2)
	
	for{
		select{// makes the execution wait for a goroutine to receive message 
		case msg,open:= <-ch1:
			if !open{
				ch1=nil    //Avoid reading from the closed channel 
			}else{
				fmt.Println(msg)

			}
		case msg,open:=<-ch2:
			if !open{
				ch2=nil
			}else{
				fmt.Println(msg)
			}
		}
		if ch1==nil && ch2==nil{
			break
		}
	}
	fmt.Println("All messages received.")
}


//Utility of this appraoch
//1.Allows handling multiple senders efficiently
//2.Ensures messages are received as soon as they arrive
//3.Avoids unnecessary blocking or waiting