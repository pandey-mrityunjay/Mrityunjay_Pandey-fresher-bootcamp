package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*	Write a program that starts with a single bank account with a starting balance of Rs.500.
	It should be possible to deposit and withdraw money concurrently.
	However, the balance update should happen only once at any point in time.
	If the withdrawal action should fail if the balance is reaching negative.

	Expectations: Use Goroutines, Mutex
*/

// Bank account with Balance and a Mutex lock, for write action at a time on the balance of the account
type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

// Deposit Function (thread-safe)
func (acc *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mutex.Lock() //Lock the balance update

	acc.balance += amount
	fmt.Printf("Amount Deposited: %d \n Account Balance Updated: %d\n", amount, acc.balance)

	acc.mutex.Unlock() //Unlock after update
}

func (acc *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mutex.Lock()

	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Printf("Withdrew Rs.%d, New Balance: Rs.%d \n", amount, acc.balance)

	}
	acc.mutex.Unlock()
}

func main() {
	account := &BankAccount{balance: 500}

	var wg sync.WaitGroup

	//Simulate concurrent transactions
	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func() {
			amount := rand.Intn(200) + 1

			if rand.Intn(2) == 0 {
				account.Deposit(amount, &wg)

			} else {
				account.Withdraw(amount, &wg)
			}
		}()

	}
	wg.Wait()

	fmt.Printf("Final Bank Account Balance: Rs%d", account.balance)

}
