package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/internal/car/model"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	
)


func main() {
	reader := bufio.NewReader(os.Stdin)

	//Get Car Brand
	brand := util.ReadString(reader, "Enter brand", "Aston Martin")
	//Get Car Model
	modelName := util.ReadString(reader, "Enter model", "DB12")
	//Get Car Speed
	speed := util.ReadInt(reader, "Enter Speed", 80)

	car := model.NewCar(brand, modelName)
	fmt.Println("Before Accelerate")
	carString := car.DisplayStatus()
	fmt.Println(carString)
	car.Accelerate(speed)
	fmt.Println("After Accelerate")
	carString = car.DisplayStatus()
	fmt.Println(carString)

}
