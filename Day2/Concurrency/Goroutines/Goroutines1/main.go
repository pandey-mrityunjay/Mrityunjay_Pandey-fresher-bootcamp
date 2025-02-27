package main

import (
	"fmt"
	"time"
)
//Concurrency=>More than one task at once =>Delayed Response Time => Interleaving operation of process


func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {

	//1.Concept here in the channel lies is to create a communication that sends the data back to the main function
	//2.Only after the execution of the commands the data is sent back to the main
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true // <- the arrow here defines the flow of data from one to another
	//close(doneChan)
}

//We can make the function calls as goroutines by adding the keyword go infront of the functions
//it tells the compiler to run these calls parallel

func main() {
	done := make(chan bool)
	go greet("Nice to meet you", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're doing well", done)
	for doneChan := range done {
		fmt.Println(doneChan)
	}

}

func done_channels() {
	//Here the main functions only task is to dispatch the goroutines
	//As soon the the goroutines are dispatched
	//The main function exits resulting in no display of output
	//can be handled using channels
	done := make(chan bool)
	go greet("Nice to meet you", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're doing well", done)
	<-done // just states that statements intends to recieve data from the <-data channel
	<-done
	<-done
	<-done
}

func done_channels_slice() {
	dones := make([]chan bool, 4)

	dones[0] = make(chan bool)
	go greet("Nice to meet you", dones[0])

	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])

	dones[3] = make(chan bool)
	go greet("I hope you're doing well", dones[3])

}
