package main

import "fmt"

func main() {
	// for range only copy the value so this wont work
	accounts := makeAccounts()
	for _, v := range accounts {
		v.Balance += 10
	}
	fmt.Println(accounts)

	// we can use index to fix that problem
	accounts = makeAccounts()
	for i := range accounts {
		accounts[i].Balance += 10
	}
	fmt.Println(accounts)

	// or use an pointer, but we can't always store a slice of pointers
	accountsPtr := makeAccountsPtr()
	for _, v := range accountsPtr {
		v.Balance += 10
	}

	for _, v := range accountsPtr {
		fmt.Println(*v)
	}
}

type Account struct {
	Id      string
	Balance int
}

func makeAccounts() []Account {
	return []Account{
		Account{"1", 1},
		Account{"2", 2},
		Account{"3", 3},
	}
}

func makeAccountsPtr() []*Account {
	return []*Account{
		&Account{"1", 1},
		&Account{"2", 2},
		&Account{"3", 3},
	}
}
