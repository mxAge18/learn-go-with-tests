package pointer

import (
	"errors"
	"fmt"
)
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}
func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposite(amount Bitcoin)  {
	// fmt.Println("address of balance in function is", &w.balance)
	w.balance += amount
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
        return InsufficientFundsError
    }
	w.balance -= amount
	return nil
}