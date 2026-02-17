package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DHRUVCHARNE/low-level-design/internal/car/model"
)

func readString(reader *bufio.Reader, prompt, defaultVal string) string {
	fmt.Printf("%s [%s] : ", prompt, defaultVal)
	input, _:= reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultVal
	}
	return input
}

func readInt(reader *bufio.Reader, prompt string, defaultVal int) int {
	fmt.Printf("%s [%d] : ", prompt, defaultVal)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == ""  {
		return defaultVal
	}
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Number, Using Default")
		return defaultVal
	}
	return val
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Get Car Brand
	brand := readString(reader, "Enter brand", "Aston Martin")
	//Get Car Model
	modelName := readString(reader, "Enter model", "DB12")
	//Get Car Speed
	speed := readInt(reader, "Enter Speed", 80)

	car := model.NewCar(brand, modelName)
	fmt.Println("Before Accelerate")
	carString := car.DisplayStatus()
	fmt.Println(carString)
	car.Accelerate(speed)
	fmt.Println("After Accelerate")
	carString = car.DisplayStatus()
	fmt.Println(carString)

}
