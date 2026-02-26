package model

import "fmt"

type TrafficLight int

const  (
	RED TrafficLight = iota
	GREEN
	YELLOW
)

func (t *TrafficLight) GetDuration() int {
	switch *t {
	case RED:
		return 30
	case GREEN:
		return 25
	case YELLOW:
		return 5
	default:
		return 0
	}
}

func (t *TrafficLight) Next(){
	switch *t {
	case RED:
		*t=GREEN
	case GREEN:
		*t = YELLOW
	case YELLOW:
		*t= RED
	}
}

func Display(light TrafficLight) {
	switch light {
	case RED:
		fmt.Printf("RED (%d)\n",light.GetDuration())
	case GREEN:
		fmt.Printf("GREEN (%d)\n",light.GetDuration())
	case YELLOW:
		fmt.Printf("YELLOW (%d)\n",light.GetDuration())
	default:
		fmt.Println("NO COLOR")
	}
}