package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	Name    string
	lock    sync.RWMutex
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance -= amount

}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

func (a *Account) GetBalance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.balance
}

func main() {
	var account Account
	var wg sync.WaitGroup

	account.Name = "Test account"

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go account.Deposit(100, &wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go account.Withdraw(100, &wg)
	}

	wg.Wait()
	fmt.Printf("Balance: %d\n", account.GetBalance())
}
