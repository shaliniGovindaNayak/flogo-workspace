package publishmqtt

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	dht "github.com/d2r2/go-dht"
	logger "github.com/d2r2/go-logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
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
type data struct {
	temp  string `json:"Temperature"`
	humid string `json:"Humidity"`
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

	defer logger.FinalizeLogger()

	topic := context.GetInput("topic").(string)
	host := context.GetInput("host").(string)
	password := context.GetInput("password").(string)
	username := context.GetInput("username").(string)
	//payload := context.GetInput("payload").(string)

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
loop:
	for {
		fmt.Println("*************************************Sample Publisher Started************************************************")

		fmt.Println("************************************doing publish*******************************************************")

		for payload := range generate() {

			client.Publish(c1.topic, byte(c1.qos), false, payload)
			//token.Wait()
			lg.Infof("Published message %s", payload)
			lg.Infof("done...")
			context.SetOutput("output", "done")
			continue loop
		}
	}

	//client.Disconnect(250)
	//fmt.Println("Sample Publisher Disconnected")

	//return
}

func generate() <-chan string {
	c := make(chan string)
	go func() {
		temperature, humidity, retried, err :=
			dht.ReadDHTxxWithRetry(dht.DHT11, 17, false, 10)
		if err != nil {
			lg.Fatal(err)
		}
		lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%% (retried %d times)",
			dht.DHT11, temperature, humidity, retried)

		c <- fmt.Sprintf(`{"temperature": "%v", "humidity": "%v"}`, temperature, humidity)

	}()

	return c
}
