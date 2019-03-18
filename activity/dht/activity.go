package dht

import (
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	pin := context.GetInput("pin").(string)
	adaptor := raspi.NewAdaptor()
	MQ := gpio.NewDirectPinDriver(adaptor, pin)

	work := func() {
		gobot.Every(1*time.Second, func() {
			value, _ := MQ.DigitalRead()
			context.SetOutput("output", value)
			fmt.Println(value)
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{MQ},
		work,
	)

	robot.Start()
	return true, nil
}
