package main

//Worker pools allows you to efficiently execute the task by dividing the task among the resoucres present, i.e the goroutines

//Advantages
//1. Faster execution
//2. Efficient resources
//3. Better load balancing

//Working of worker pools
// Create "Task Queue"=>  Start Multiple "Workers"(Goroutines)  => Collect Results in a "Results Channel"=> Wait for all workers to finsih

import (
	"fmt"
	"sync"
	"time"
)

// Worker function(each worker does one task at a time)
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // mark this worker as done after finishing all task

	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d \n", id, job)
		time.Sleep(time.Second) //Simulate work

		results <- fmt.Sprintf("Job %d done by worker %d", job, id)

	}
}

func main() {
	const numWorker = 3
	const numJobs = 10

	jobs := make(chan int, numJobs) //Channel for tasks

	results := make(chan string, numJobs) //Channel for results

	var wg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)

	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait() //wait for all workers to finish

	close(results)

	for result := range results {
		fmt.Println(result)
	}
	fmt.Println("All jobs completed")

}
