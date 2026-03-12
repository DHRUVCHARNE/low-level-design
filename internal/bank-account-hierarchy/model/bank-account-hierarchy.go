package model

import "fmt"

const MINIMUM_BALANCE = 100
type Account interface {
	Deposit(amount float64) bool
	Withdraw(amount float64) bool
	DisplayAccount()
    GetAccountNumber() string
}
type BankAccount struct {
	ownerName     string
	accountNumber string
	balance       float64
}

func NewBankAccount(ownerName, accountNumber string, balance float64) *BankAccount {
	return &BankAccount{ownerName: ownerName, accountNumber: accountNumber, balance: balance}
}

func (b *BankAccount) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	b.balance = b.balance + amount
	return true
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount > b.balance {
		return false
	}
	b.balance = b.balance - amount
	return true
}

func (b *BankAccount) DisplayAccount() {
	fmt.Printf("%s (%s) | Balance: %.2f", b.ownerName, b.accountNumber, b.balance)
}

func (b *BankAccount) GetAccountNumber() string {
	return b.accountNumber
}

type SavingsAccount struct {
	BankAccount
	interestRate float64
}

func NewSavingsAccount(ownerName, accountNumber string, balance, interestRate float64) *SavingsAccount {
	return &SavingsAccount{BankAccount: BankAccount{ownerName: ownerName, accountNumber: accountNumber, balance: balance}, interestRate: interestRate}
}

func (b *SavingsAccount) Withdraw(amount float64) bool {
	if (b.BankAccount.balance - amount) >= MINIMUM_BALANCE {
		b.BankAccount.balance = b.BankAccount.balance - amount
		return true
	}
	return false
}

func (b *SavingsAccount) ApplyInterest() {
	b.BankAccount.balance = b.BankAccount.balance * (100 + b.interestRate) / 100
}

type CheckingAccount struct {
	overdraftLimit float64
	BankAccount      
}

func NewCheckingAccount(ownerName, accountNumber string, balance, overdraftLimit float64) *CheckingAccount {
	return &CheckingAccount{BankAccount: BankAccount{ownerName: ownerName, accountNumber: accountNumber, balance: balance}, overdraftLimit: overdraftLimit}
}

func (b *CheckingAccount) Withdraw(amount float64) bool {
	if (b.BankAccount.balance + b.overdraftLimit) >= amount {
		b.BankAccount.balance = b.BankAccount.balance - amount
		return true
	}
	return false
}
