package main


import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/bank-account-hierarchy/model"
)

func main() {
	fmt.Println("Multiple Bank Account CLI")

	reader := bufio.NewReader(os.Stdin)

	accounts := map[string]model.Account{}

	for {
		choice := util.ReadByte(
			reader,
			"\nEnter 'c' to create a bank account"+
				"\nEnter 'a' to list all bank accounts"+
				"\nEnter 'd' to deposit"+
				"\nEnter 'w' to withdraw"+
				"\nEnter any other key to quit\n",
			'q',
		)

		switch choice {

		case 'c':
			acc, err := CreateBankAccount(reader, &accounts)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if acc != nil {
				accounts[acc.GetAccountNumber()] = acc
				fmt.Println("Account Created Successfully")
			}

		case 'a':
			if len(accounts) == 0 {
				fmt.Println("No accounts found")
				continue
			}

			fmt.Println("\nAccounts:")
			for _, acc := range accounts {
				acc.DisplayAccount()
				fmt.Println()
			}

		case 'd':
			accNumber := util.ReadString(reader, "Enter account number", "")

			acc, ok := accounts[accNumber]
			if !ok {
				fmt.Println("Account not found")
				continue
			}

			amount := util.ReadFloat64(reader, "Enter deposit amount", 0)

			if acc.Deposit(amount) {
				fmt.Println("Deposit successful")
			} else {
				fmt.Println("Deposit failed")
			}

		case 'w':
			accNumber := util.ReadString(reader, "Enter account number", "")

			acc, ok := accounts[accNumber]
			if !ok {
				fmt.Println("Account not found")
				continue
			}

			amount := util.ReadFloat64(reader, "Enter withdrawal amount", 0)

			if acc.Withdraw(amount) {
				fmt.Println("Withdrawal successful")
			} else {
				fmt.Println("Withdrawal failed")
			}

		default:
			fmt.Println("Exiting the CLI")
			return
		}
	}
}

func CreateBankAccount(reader *bufio.Reader, accounts *map[string]model.Account) (model.Account, error) {

	choice := util.ReadByte(
		reader,
		"Enter 'a' to create a Bank Account\n"+
			"Enter 'b' to create a Savings Account\n"+
			"Enter 'c' to create a Checking Account\n"+
			"Any other key to exit\n",
		'a',
	)

	switch choice {

	case 'a':
		accountNumber, ownerName, balance, err := ReadAccountInfo(reader, len(*accounts))
		if err != nil {
			return nil, err
		}

		return model.NewBankAccount(ownerName, accountNumber, balance), nil

	case 'b':
		accountNumber, ownerName, balance, err := ReadAccountInfo(reader, len(*accounts))
		if err != nil {
			return nil, err
		}

		interestRate := util.ReadFloat64(reader, "Enter the Interest Rate", 5)

		if interestRate < 0 || interestRate > 10 {
			return nil, fmt.Errorf("interest rate out of range")
		}

		return model.NewSavingsAccount(ownerName, accountNumber, balance, interestRate), nil

	case 'c':
		accountNumber, ownerName, balance, err := ReadAccountInfo(reader, len(*accounts))
		if err != nil {
			return nil, err
		}

		overdraftLimit := util.ReadFloat64(reader, "Enter overdraft limit", 5)

		if overdraftLimit < 0 {
			return nil, fmt.Errorf("overdraft cannot be negative")
		}

		return model.NewCheckingAccount(ownerName, accountNumber, balance, overdraftLimit), nil

	default:
		return nil, nil
	}
}

func ReadAccountInfo(reader *bufio.Reader,count int) (string,string,float64,error) {
	accountNumber:=util.ReadString(reader,"Enter the Account Number",fmt.Sprintf("Acc-%d",count))
		ownerName:=util.ReadString(reader,"Enter the Owner Name",fmt.Sprintf("OW-%d",count))
		balance:=util.ReadFloat64(reader,"Enter the Balance",0.0)
		if balance < 0 {
			return "","",0,fmt.Errorf("balance cannot be negative")
		}
		return  accountNumber,ownerName,balance,nil
}