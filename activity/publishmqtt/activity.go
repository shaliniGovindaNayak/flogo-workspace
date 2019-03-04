package publishmqtt

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type params struct {
	topic     string
	broker    string
	password  string
	user      string
	id        string
	cleansess bool
	qos       int
	num       int

	action string
	store  string
}

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

	topic := context.GetInput("topic").(string)
	host := context.GetInput("host").(string)
	password := context.GetInput("password").(string)
	username := context.GetInput("username").(string)
	payload := context.GetInput("payload").(string)

	credentails := params{topic, host, password, username, "host", false, 0, 1, "pub", ":memory"}
	c1 := &credentails

	if c1.topic == "" {
		fmt.Println("Invalid setting for -topic, must not be empty")
		return
	}

	fmt.Printf("Sample Info:\n")
	fmt.Printf("\taction:    %s\n", c1.action)
	fmt.Printf("\tbroker:    %s\n", c1.broker)
	fmt.Printf("\tclientid:  %s\n", c1.id)
	fmt.Printf("\tuser:      %s\n", c1.user)
	fmt.Printf("\tpassword:  %s\n", c1.password)
	fmt.Printf("\ttopic:     %s\n", c1.topic)
	fmt.Printf("\tqos:       %d\n", c1.qos)
	fmt.Printf("\tcleansess: %v\n", c1.cleansess)
	fmt.Printf("\tnum:       %d\n", c1.num)
	fmt.Printf("\tstore:     %s\n", c1.store)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(c1.broker)
	opts.SetClientID(c1.id)
	opts.SetUsername(c1.user)
	opts.SetPassword(c1.password)
	opts.SetCleanSession(c1.cleansess)
	if c1.store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(c1.store))
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Publish(c1.topic, byte(c1.qos), false, payload)
	context.SetOutput("output", "Done..")
	return true, nil
}
