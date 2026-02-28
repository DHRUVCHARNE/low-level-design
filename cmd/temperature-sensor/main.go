package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/temperature-sensor/model"
)

func main() {
	fmt.Println("Temperature Sensors CLI")
	reader := bufio.NewReader(os.Stdin)
	sensor:=model.NewTemperatureSensor()
    
	for {
	choice:=util.ReadByte(reader,"Enter 'a' for adding temperature\nEnter 'c' for getting the count\nEnter 'm' for getting the mean temperature\nEnter 'q' for exiting the CLI\n",'a')
	switch choice {
	case 'a':
		temp:=util.ReadFloat64(reader,"Enter the Temperature in Celcius",35.7)
		sensor.AddReading(temp)
		fmt.Println("Added to sensor")
	case 'c':
		fmt.Printf("The Count of the Temperatures in the Sensor is: %.3f\n",sensor.GetReadings())
	case 'm':
		fmt.Printf("The Average of the Temperatures in the Sensor is: %.3f\n",sensor.GetAverage())
    case 'q':
		fmt.Println("Exiting the CLI")
		return
	}
	}
}