package main

//A non-blocking channel operation allows a goroutine to attempt sending or recieving from a channel without waiting
//Server Monitoring System
//1.Process messages from active servers immediately.
//2.Skip waiting if no message is available and perform other tasks.
//3.Prevent blocking in case a channel is not ready

import(
	"fmt"
	"time"
)
func main(){
	ch:=make(chan string)
	go func(){
		time.Sleep(3*time.Second)
		ch <- "Server is healthy"
	}()
	for i:=1;i<=5;i++{
		select{
		case msg:=<-ch:
			fmt.Println("Recieved:",msg)
		default:
			fmt.Println("No message yet,checking system logs...")
		}
		time.Sleep(1*time.Second)
	}
	fmt.Println("Finished Monitoring")
}
