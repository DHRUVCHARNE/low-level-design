package model

import "fmt"

// Define a Custon Type for Enum
type OrderStatus int

// Declaring related constants using iota
const (
	PLACED OrderStatus = iota
	CONFIRMED
	SHIPPED
	DELIVERED
	CANCELLED
)

func (s OrderStatus) orderStatusToString() string {
	switch s {
	case PLACED :
		return "PLACED"
	case CONFIRMED:
		return "CONFIRMED"
	case SHIPPED:
		return "SHIPPED"
	case DELIVERED:
		return "DELIVERED"
	case CANCELLED:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}

type PaymentMethod struct {
	displayName string
	feePercent float64
}
//group multiple variable declaration
var (
	CreditCard = PaymentMethod{
		displayName: "Credit Card",
		feePercent:  2.5,
	}

	DebitCard = PaymentMethod{
		displayName: "Debit Card",
		feePercent:  1.0,
	}

	UPI = PaymentMethod{
		displayName: "UPI",
		feePercent:  0.0,
	}

	NetBanking = PaymentMethod{
		displayName: "Net Banking",
		feePercent:  1.5,
	}
)

type Order struct {
	orderId string
	orderStatus OrderStatus
	paymentMethod PaymentMethod
	amount float64
}

func NewOrder(orderId string,paymentMethod PaymentMethod,amount float64) (*Order) {
	return &Order{orderId: orderId,paymentMethod: paymentMethod,amount: amount}
}

func (o *Order) AdvanceStatus() bool {
	switch o.orderStatus {
	case PLACED:
		o.orderStatus=CONFIRMED
		return true
	case CONFIRMED:
		o.orderStatus=SHIPPED
		return true
	case SHIPPED:
		o.orderStatus=DELIVERED
		return true
	default:
		return false
	}
}

func (o *Order) Cancel() bool {
	if o.orderStatus == PLACED || o.orderStatus == CONFIRMED {
		o.orderStatus = CANCELLED
		return true
	}
	return false
}

func (o *Order) GetTotalWithFees() float64 {
	return o.amount*(100+o.paymentMethod.feePercent)/100
}

func (o *Order) DisplayInfo() {
	fmt.Printf("Order %s | Status: %s | Payment: %s | Amount: $%.2f (with fees: $%.2f)\n",o.orderId,o.orderStatus.orderStatusToString(),o.paymentMethod.displayName,o.amount,o.GetTotalWithFees())
}