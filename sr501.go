package sr501_mod

import (
	"periph.io/x/conn/v3/gpio"
)

type Sensor struct {
	pinName string
	pin     gpio.PinIO
	Events  chan bool
}
