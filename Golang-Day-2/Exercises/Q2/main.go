package main

/*Suppose you are a class monitor and you have to take the performance rating of the class teacher from your classmates. Write a program to take the ratings from all of your classmates and then print the average rating.

Expectation: use waitgroups

Assumptions
There are 200 students in the class.
Every student will take a random amount of time to respond.
You can simulate the random response time of your classmates by using the math/rand package.
*/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to simulate a student giving a rating

func giveRating(id int, ratings chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	rating := rand.Intn(5) + 1

	fmt.Printf("Student %d gave a rating: %d \n", id, rating)

	ratings <- rating

}

func main() {
	rand.Seed(time.Now().UnixNano())

	totalStudents := 200
	ratings := make(chan int, totalStudents)

	var wg sync.WaitGroup //WaitGroup to synchronize goroutines

	//Simulate 200 students
	for i := 1; i <= totalStudents; i++ {
		wg.Add(1)
		go giveRating(i, ratings, &wg)
	}

	//Closing all the channel once all ratings are collected
	go func() {
		wg.Wait()
		close(ratings)
	}()

	var sum, count int

	//calculate and compute the average rating
	for rating := range ratings {
		sum += rating
		count++
	}

	averageRating := float64(sum) / float64(count)

	fmt.Printf("\n Total Students: %d \n", count)
	fmt.Printf("Average Rating: %.2f \n", averageRating)

}
