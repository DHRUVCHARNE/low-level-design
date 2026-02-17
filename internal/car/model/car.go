package model

import "fmt"

type Car struct {
	brand string
	model string
	speed int 
}
//Function
func NewCar(brand, model string) *Car {
	return &Car{brand:brand,model:model}
}
//Method 
func (c *Car) Accelerate(increment int){
	c.speed+=increment
}

func (c *Car) Speed() int {
	return c.speed
}

func (c *Car) Model() string {
	return c.model
}

func(c *Car) Brand() string {
	return c.brand
}

func (c *Car) DisplayStatus() string {
	return fmt.Sprintf(
		"%s %s is running at %d km/h",
		c.brand, c.model, c.speed,
	)
}