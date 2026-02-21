package model

import "fmt"

type FoodOrder struct {
	orderId      string
	customerName string
	items        []string
	totalAmount  float64
	isPlaced     bool
}

func NewFoodOrder(orderId, customerName string) *FoodOrder {
	return &FoodOrder{orderId: orderId, customerName: customerName}
}

func (f *FoodOrder) AddItem(name string, price float64) {
	if(f.isPlaced){
		fmt.Println("Cannot modify a placed order")
		return
	}
	f.items=append(f.items, name)
	f.totalAmount+=price
}

func (f *FoodOrder) PlaceOrder() bool {
	if(f.isPlaced || len(f.items)==0){
		return false
	}
	f.isPlaced=true
	return true
}

func (f *FoodOrder) GetItemCount() int{
	return len(f.items)
}

func (f *FoodOrder) DisplayOrder() {
	status:="PENDING"
	if f.isPlaced {
		status="PLACED"
	}
	fmt.Println("Order ",f.orderId,"(",f.customerName,") - ",status)
	for _,item := range f.items {
    fmt.Println(" - ",item)
	}
	fmt.Printf("Total: $ %.2f\n",f.totalAmount)
}