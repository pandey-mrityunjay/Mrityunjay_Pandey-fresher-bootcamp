package main

/*1. For a given set of strings, find out the frequency of each letter in all the strings parallelly.
For example, if given the following input.,
Expectations: Slices, Goroutines, Channels
[“quick”
,
“brown”
,
“fox”
,
“lazy”
,
“dog”]
The output should be.,
{
“a”: 1,
“b”: 1,
“c”: 1,
…
“z”: 1
2. }  */

import (
	"bufio"            
	"fmt"
	"os"
	"strings"
	"sync"
)

// function to process each string and send frequency counts to a channel
func countLetters(s string, ch chan map[rune]int, wg *sync.WaitGroup) { //map[rune] instread of map[byte] as it supports all unicode characters
	defer wg.Done()
	freq := make(map[rune]int)

	for _, char := range s {
		if char != ' ' {
			freq[char]++
		}
	}

	ch <- freq

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)    //creates a scanner to read user input
	fmt.Println("Enter words separated by spaces:")

	scanner.Scan()
	input := scanner.Text()            //reads the user input(one line of the text)

	words := strings.Fields(input)     //splits into words using spaces
									   //converts the input string into a slice of words

	ch := make(chan map[rune]int, len(words))  //Bufferend Channel to hold frequency maps

	var wg sync.WaitGroup				//Creates a WaitGroup to synchronize goroutines

	for _, word := range words {
		wg.Add(1)
		go countLetters(word, ch, &wg)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	finalFreq := make(map[rune]int)

	for freqMap := range ch {
		for char, count := range freqMap {
			finalFreq[char] += count
		}
	}
	fmt.Println("Letter Frequency:")
	for char, count := range finalFreq {
		fmt.Printf("%q: %d \n", char, count)
	}
}


/*
Slices     - To store words from the user input
Goroutines - Each word is processed in parallel
Channels   - Used to pass frequency maps between goroutines and main function
WaitGroup  - Ensures all goroutines complete before closing the channel
Concurrency- Achieves faster execution by processing words in parallel


*/