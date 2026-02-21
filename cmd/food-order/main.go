package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/food-order/model"
)
// readItem reads one item.
// Returns (name, price, done)
func readItem(reader *bufio.Reader) (string, float64, bool) {
	name := util.ReadString(
		reader,
		"Enter Item Name (type 'done' to finish, 'exit' to quit)",
		"Pizza",
	)

	if name == "done" {
		return "", 0, true
	}

	if name == "exit" {
		fmt.Println("\n👋 Exiting program")
		os.Exit(0)
	}

	price := util.ReadFloat64(reader, "Enter Item Price", 5.75)
	return name, price, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🍕 Food Ordering CLI")
	fmt.Println("Type 'done' to finish items, 'exit' to quit")

	orderCounter := 1

	for {
		fmt.Printf("\n--- Order No. %d ---\n", orderCounter)

		customerName := util.ReadString(
			reader,
			"Enter Customer Name",
			fmt.Sprintf("CUS-%d", orderCounter),
		)

		orderID := util.ReadString(
			reader,
			"Enter Order ID",
			fmt.Sprintf("ORD-%d", orderCounter),
		)

		order := model.NewFoodOrder(orderID, customerName)

		fmt.Println("\nAdd Items")

		for {
			name, price, done := readItem(reader)
			if done {
				fmt.Println("🛑 Finished adding items")
				break
			}
			order.AddItem(name, price)
		}

		if order.PlaceOrder() {
			fmt.Println("\n✅ Order Placed Successfully")
		} else {
			fmt.Println("\n❌ Failed to Place Order")
		}

		order.DisplayOrder()
		orderCounter++

		cont := util.ReadString(
			reader,
			"\nPress Enter to place another order or type 'exit'",
			"",
		)

		if cont == "exit" {
			fmt.Println("\n👋 Program exited cleanly")
			return
		}
	}
}