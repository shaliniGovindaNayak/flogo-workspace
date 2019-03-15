package dht

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	dht11 "github.com/d2r2/go-dht"
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

	sensorType := context.GetInput("sensorType").(string)
	gpiopin := context.GetInput("gpiopin").(int)

	sensor := dht11.DHT11

	if sensorType == "dht22" {
		sensor = dht11.DHT22
	}

	temperature, humidity, err := dht11.ReadDHTxx(sensor, gpiopin, false)

	context.SetOutput("temperature", temperature)
	context.SetOutput("humidity", humidity)

	return true, nil
}
