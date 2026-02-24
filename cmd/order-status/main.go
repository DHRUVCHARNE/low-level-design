package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/order-status/model"
)

func main() {
	fmt.Println("Order Status CLI")
	reader := bufio.NewReader(os.Stdin)

	var orders []*model.Order
	orderCounter := 1

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1) Create new order")
		fmt.Println("2) List orders")
		fmt.Println("3) Advance order status")
		fmt.Println("4) Cancel order")
		fmt.Println("5) Show order details")
		fmt.Println("6) Exit")

		choice := util.ReadInt(reader, "Select option", 1)

		switch choice {
		case 1:
			createOrder(reader, &orders, &orderCounter)
		case 2:
			listOrders(orders)
		case 3:
			advanceOrder(reader, orders)
		case 4:
			cancelOrder(reader, orders)
		case 5:
			showOrderDetails(reader, orders)
		case 6:
			fmt.Println("\n👋 Exiting. Bye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func createOrder(reader *bufio.Reader, orders *[]*model.Order, counter *int) {
	fmt.Printf("\n--- Create Order #%d ---\n", *counter)

	customer := util.ReadString(reader, "Enter Customer Name", fmt.Sprintf("CUS-%d", *counter))
	orderID := util.ReadString(reader, "Enter Order ID", fmt.Sprintf("ORD-%d", *counter))

	// choose payment method
	fmt.Println("\nChoose Payment Method:")
	fmt.Println("1) Credit Card")
	fmt.Println("2) Debit Card")
	fmt.Println("3) UPI")
	fmt.Println("4) Net Banking")

	pmChoice := util.ReadInt(reader, "Payment option", 1)
	var pm model.PaymentMethod
	switch pmChoice {
	case 1:
		pm = model.CreditCard
	case 2:
		pm = model.DebitCard
	case 3:
		pm = model.UPI
	case 4:
		pm = model.NetBanking
	default:
		fmt.Println("Invalid choice, defaulting to UPI")
		pm = model.UPI
	}

	amount := util.ReadFloat64(reader, "Enter Order Amount", 0.0)

	o := model.NewOrder(orderID, pm, amount)
	// note: NewOrder doesn't set customerName (your model doesn't store it); if you want to include it,
	// add a field to model.Order and set it here.
	fmt.Printf("\n✅ Created order %s for %s\n", orderID, customer)
	o.DisplayInfo()

	*orders = append(*orders, o)
	*counter++
}

func listOrders(orders []*model.Order) {
	fmt.Println("\n--- Orders ---")
	if len(orders) == 0 {
		fmt.Println("No orders yet")
		return
	}
	for i, o := range orders {
		fmt.Printf("%d) ", i+1)
		o.DisplayInfo()
	}
}

func findOrderByID(orders []*model.Order, id string) (*model.Order, int) {
	for i, o := range orders {
		// the orderId field is unexported; but DisplayInfo shows the id.
		// Since the struct fields are unexported, we cannot read them directly here.
		// We'll compare by rendering the DisplayInfo string — but better is to extend the model to expose an accessor.
		// However we can try matching using fmt.Sprintf of DisplayInfo output — simpler approach:
		// Instead, rely on the user choosing by list index in other flows.
		// For this helper we'll assume the orderId is in DisplayInfo output; so this function is not reliable.
		// Safer approach: return nil (so we use index-based selection in the other functions).
		_ = o
		_ = i
	}
	return nil, -1
}

func selectOrderByIndex(reader *bufio.Reader, orders []*model.Order) (*model.Order, bool) {
	if len(orders) == 0 {
		fmt.Println("No orders available")
		return nil, false
	}
	listOrders(orders)
	idx := util.ReadInt(reader, "Enter order number", 1)
	if idx < 1 || idx > len(orders) {
		fmt.Println("Invalid order number")
		return nil, false
	}
	return orders[idx-1], true
}

func advanceOrder(reader *bufio.Reader, orders []*model.Order) {
	fmt.Println("\n--- Advance Order Status ---")
	o, ok := selectOrderByIndex(reader, orders)
	if !ok {
		return
	}
	if o.AdvanceStatus() {
		fmt.Println("✅ Status advanced")
	} else {
		fmt.Println("❌ Cannot advance status further")
	}
	o.DisplayInfo()
}

func cancelOrder(reader *bufio.Reader, orders []*model.Order) {
	fmt.Println("\n--- Cancel Order ---")
	o, ok := selectOrderByIndex(reader, orders)
	if !ok {
		return
	}
	if o.Cancel() {
		fmt.Println("✅ Order cancelled")
	} else {
		fmt.Println("❌ Order cannot be cancelled at this stage")
	}
	o.DisplayInfo()
}

func showOrderDetails(reader *bufio.Reader, orders []*model.Order) {
	fmt.Println("\n--- Order Details ---")
	o, ok := selectOrderByIndex(reader, orders)
	if !ok {
		return
	}
	o.DisplayInfo()

	// Allow user to perform quick actions
	fmt.Println("\nActions: [A]dvance  [C]ancel  [Enter] back to menu")
	choice := util.ReadString(reader, "Choose action", "")
	choice = strings.ToLower(strings.TrimSpace(choice))
	switch choice {
	case "a":
		if o.AdvanceStatus() {
			fmt.Println("✅ Status advanced")
		} else {
			fmt.Println("❌ Cannot advance status")
		}
		o.DisplayInfo()
	case "c":
		if o.Cancel() {
			fmt.Println("✅ Order cancelled")
		} else {
			fmt.Println("❌ Cannot cancel order")
		}
		o.DisplayInfo()
	default:
		// back to menu
	}
}