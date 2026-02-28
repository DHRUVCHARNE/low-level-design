package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/shopping-cart/model"
)

func main() {
	fmt.Println("Shopping Cart CLI")
	reader := bufio.NewReader(os.Stdin)
	cart:=model.NewShoppingCart()
	for {
	choice:=util.ReadByte(reader,"Enter 'a' for adding Item\nEnter 'c' for getting the total\nEnter 'm' for getting the discount\nEnter 'o' for checking out\nEnter 'd' for displaying the cart\nEnter 'q' for exiting the CLI\n",'a')
	switch choice {
	case 'a':
		name:=util.ReadString(reader,"Enter the Item name","ouva Notebook")
		price:=util.ReadFloat64(reader,"Enter the price",40.00)
		cart.AddItem(name,price)
	case 'c':
		fmt.Printf("The Total of the Shopping Cart is: %.3f\n",cart.GetTotal())
	case 'm':
		code:=util.ReadString(reader,"Enter the coupon code","SAVE10")
		cart.ApplyDiscount(code)
	case 'o':
		cart.Checkout()
	case 'd':
		cart.Display()
    case 'q':
		fmt.Println("Exiting the CLI")
		return
	}
	}
}