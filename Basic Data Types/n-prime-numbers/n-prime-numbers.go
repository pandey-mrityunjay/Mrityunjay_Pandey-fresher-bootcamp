package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count, num := 0, 2
	for count < 10 {
		if isPrime(num) {
			fmt.Println(num)
			count++
		}
		num++
	}
}
