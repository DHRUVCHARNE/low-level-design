package model

type TemperatureSensor struct {
	readings [] float64
}

func NewTemperatureSensor() *TemperatureSensor {
	return &TemperatureSensor{}
}

func (t *TemperatureSensor) AddReading(value float64) {
	t.readings = append(t.readings,value)
}

func (t *TemperatureSensor) GetAverage() float64 {
	sum:=0.0
	if len(t.readings) == 0 {
    return sum
	}
	for _,v := range t.readings {
		sum+=v
	}
	return sum/float64(len(t.readings))
}