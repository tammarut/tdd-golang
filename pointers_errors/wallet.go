package wallet

import "fmt"

import "errors"

type Stringer interface {
	String() string
}

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
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("Oh no")
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
