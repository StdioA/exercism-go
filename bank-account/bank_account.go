package account

import "sync"


type Account struct {
	balance int64
	opening bool
	sync.RWMutex
}


func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, opening: true}
}


func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.opening {
		return 0, false
	}

	account.opening = false
	return account.balance, true
}


func (account *Account) Balance() (balance int64, ok bool) {
	account.RLock()
	defer account.RUnlock()
	if !account.opening {
		return 0, false
	}

	return account.balance, true
}


func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.opening {
		return 0, false
	}
	balance := account.balance + amount
	if balance < 0 {
		return account.balance, false
	}
	account.balance = balance
	return balance, true
}
