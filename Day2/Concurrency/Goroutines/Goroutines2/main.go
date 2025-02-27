package main

//Goroutines are the functions or methods that run concurrently with other functions. They are managed by the Go runtime
//rather than the OS

//A goroutine exits if the main function exits
// If main() finishes execution before a goroutine the goroutine is terminated

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello,Go!"
}

func main() {
	ch := make(chan string)//creating a channel

	go sendData(ch)   //Goroutine sends

	fmt.Println(<-ch) //main goroutine receives data
}

/*
1. Channels block until data is received/sent
2. Unbuffered channels synchronize sender and receiver
*/