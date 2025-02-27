package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup helps synchronize multiple goroutines
//WaitGroup ensures goroutines complete before exiting
//Ensuring no premature program termination while goroutines are running


func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Decrements the counter
	fmt.Printf("Worker %d started \n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increment counter
		go worker(i, &wg)
	}
	wg.Wait()

	fmt.Println("All workers done")
}


