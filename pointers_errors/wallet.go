package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(coin int) {
	w.balance += coin
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
}

func (w *Wallet) Balance() int {
	return w.balance
}
