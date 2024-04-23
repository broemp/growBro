package sensortypes

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

type Sense interface{}

type ConnectionType int

type GPIOSensor struct {
	Model      string
	Connection ConnectionType
	GPIOPin    int
	Path       string
	Pin        rpio.Pin
}

func (s GPIOSensor) Init() {
	Pin := rpio.Pin(s.GPIOPin)
	Pin.Input()
}

func (s GPIOSensor) Read() {
	res := rpio.ReadPin(rpio.Pin(s.GPIOPin))
	fmt.Printf("res: %v\n", res)
}
