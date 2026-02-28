package model

import "fmt"

type ShoppingCart struct {
	items map[string]float64
	discountApplied bool
	isCheckedOut bool
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{items:make(map[string]float64),discountApplied: false,isCheckedOut: false}
}

func (s *ShoppingCart) AddItem(name string ,price float64) {
	if s.isCheckedOut {
		fmt.Println("Error: Already Checked Out")
		return
	}
	s.items[name]=price
}

func (s *ShoppingCart) ApplyDiscount(code string) bool {
	if s.discountApplied || s.isCheckedOut || code !="SAVE10" {
		return false
	}
	s.discountApplied=true
	return true
}

func (s *ShoppingCart) GetTotal() float64 {
	var sum float64=0
	for _,price:=range s.items {
    sum+=price
	}
	return sum
}

func (s *ShoppingCart) Checkout() {
	s.isCheckedOut=true
}

func (s *ShoppingCart) Display() {
	i:=1
	for name,value :=range s.items {
		fmt.Printf("%d. %s --- %.2f\n",i,name,value)
		i++
	}
	discount:="APPLIED"
	if !s.discountApplied {
		discount="NOT APPLIED"
	}
	checkedout := "CHECKED OUT"
	if ! s.isCheckedOut {
		checkedout="NOT CHECKED OUT"
	}
	fmt.Printf("DISCOUNT: %s\n",discount)
	fmt.Println(checkedout)
	fmt.Println("Total Cart Value: Rs ",s.GetTotal())

}