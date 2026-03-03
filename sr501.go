package sr501_mod

import (
	"fmt"

	"github.com/fatih/color"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Sensor struct {
	pinName string
	pin     gpio.PinIO
	Events  chan bool
}

func NewSensor(pinName string) (*Sensor, error) {
	if _, err := host.Init(); err != nil {
		color.Red("Failed to get hardware registers for HC-SR501")

		return nil, err
	}

	pin := gpioreg.ByName(pinName)

	if pin == nil {
		color.Red("Failed to get pin %s for HC-SR501", pinName)
		return nil, fmt.Errorf("Failed to get pin")
	}

	if err := pin.In(gpio.PullDown, gpio.BothEdges); err != nil {
		return nil, err
	}

	return &Sensor{
		pinName: pinName,
		pin:     pin,
		Events:  make(chan bool),
	}, nil
}

func (sensor *Sensor) Watch() {
	for {
		if sensor.pin.WaitForEdge(-1) {
			sensor.Events <- (sensor.pin.Read() == gpio.High)
		}
	}
}
