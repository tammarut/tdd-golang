package wallet

import "fmt"

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(coin Bitcoin) {
	w.balance += coin
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
