package dht11

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	dht "github.com/d2r2/go-dht"
)

const (
	ivType     = "type"
	ivPin      = "pin"
	ovTemp     = "temp"
	ovHumidity = "humidity"
)

var log = logger.GetLogger("go-dht")

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

	deviceType := context.GetInput(ivType).(string)
	gpioPin := context.GetInput(ivPin).(int)

	sensorType := dht.DHT22

	if deviceType == "DHT11" {
		sensorType = dht.DHT11
	}

	humidity, temperature, retried, err := dht.ReadDHTxxWithRetry(sensorType, gpioPin, false, 10)

	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Debugf("DHT Sensor returned [%v] temperature and [%v] humidity", temperature, humidity)
	fmt.Println(retried)
	context.SetOutput(ovTemp, fmt.Sprint(temperature))
	context.SetOutput(ovHumidity, fmt.Sprint(humidity))
	return true, nil
}
