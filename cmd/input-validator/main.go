package main


import (
	"bufio"
	"fmt"
	"os"

	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/input-validator/model"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Input Validator CLI")
	email_validator :=model.NewRegistrationService([]model.Validator{&model.EmailValidator{}})
	password_validator:=model.NewRegistrationService([]model.Validator{&model.PasswordValidator{}})
	email_and_password_validator :=model.NewRegistrationService([]model.Validator{&model.EmailValidator{},&model.PasswordValidator{}})

	for {
		choice:=util.ReadByte(reader,"For Email Validation enter e,for password validation enter p\n and for both enter b to exit enter any other key\n",'b')
		switch choice {
		case 'e':
			message:= util.ReadString(reader,"Enter the String","my_example@gmail.com")
			email_validator.Register(message)

		case 'p':
			message:= util.ReadString(reader,"Enter the String","my_example@gmail.com")
			password_validator.Register(message)
		case 'b':
			message:= util.ReadString(reader,"Enter the String","my_example@gmail.com")
			email_and_password_validator.Register(message)
		default:
			return
		}
	}
}