package main

import (
	"bufio"
	"fmt"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/bank-account/model"
	"os"
)

func main() {
	fmt.Println("Bank Account CLI")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---Creating a New Bank Account---")
	accountNumber := util.ReadString(reader, "Enter the Account Number", "Account1")
	ownerName := util.ReadString(reader, "Enter the Name of the Account Owner", "Foo")
	account := model.NewBankAccount(accountNumber, ownerName)
	fmt.Println("---New Bank Account Created---")
	fmt.Printf("---Current Balance: %.2f ---\n", account.GetBalance())
	for {
		fmt.Println("Type 'd' for deposit, 'w' for withdrawal and 'b' for balance inquiry and 'q' for exiting the cli")
		choice:=util.ReadByte(reader,"Enter the choice character",'q')
		switch choice {
		case 'd':
			amount:=util.ReadFloat64(reader,"Enter the Amount",0)
			if amount<=0{
				fmt.Println("Amount Can't be less than zero")
				continue
			}
			account.Deposit(amount)
        case 'w':
			amount:=util.ReadFloat64(reader,"Enter the Amount",0)
			if amount>= account.GetBalance() {
				fmt.Printf("You don't have Enough Balance it is %.2f and you want to withdraw %.2f\n",account.GetBalance(),amount)
				continue
			}
			account.Withdraw(amount)
		case 'b':
			fmt.Printf("Your Current Balance is: %.2f",account.GetBalance())
		case 'q':
			fmt.Println("---Exiting the Bank Account CLI---")
			return
		default:
			continue
		}
	}
}
