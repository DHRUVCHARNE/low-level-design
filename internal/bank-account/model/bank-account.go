package model

type BankAccount struct {
	accountNumber string
	ownerName string
	balance float64
}
const INITIAL_STARTING_BALANCE=0

func NewBankAccount(accountNumber,ownerName string) *BankAccount {
	return &BankAccount{accountNumber: accountNumber,ownerName: ownerName,balance: INITIAL_STARTING_BALANCE}
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance+=amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.balance-=amount
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}
