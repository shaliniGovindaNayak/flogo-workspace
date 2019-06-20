package main

import (
	"fmt"
	"time"

	dht "github.com/d2r2/go-dht"

	logger "github.com/d2r2/go-logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var lg = logger.NewPackageLogger("main", logger.DebugLevel)

//Create a structure of MQTT credentials required to publish the message
type mqttParameters struct {
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
	TEMP  string `json:"Temperature"`
	HUMID string `json:"Humidity"`
}

func main() {

}
func pub() {
	defer logger.FinalizeLogger()

	lg.Notify("***************************************************************************************************")
	lg.Notify("*** You can change verbosity of output, to modify logging level of module \"dht\"")
	lg.Notify("*** Uncomment/comment corresponding lines with call to ChangePackageLogLevel(...)")

	//MQTT credentails are assinged to variabl with respect to a specific host- 192.168.0.73(Raspberry pi)
	credentails := mqttParameters{"topic", "tcp://192.168.1.59:1883", "password", "username", "host", false, 0, 1, "pub", ":memory"}
	mqtt := &credentails

	if mqtt.topic == "" {
		fmt.Println("Invalid setting for -topic, must not be empty")
		return
	}

	fmt.Printf("Sample Info:\n")
	fmt.Printf("\taction to be performed:    %s\n", mqtt.action)
	fmt.Printf("\tbroker used:    %s\n", mqtt.broker)
	fmt.Printf("\tclientid:  %s\n", mqtt.id)
	fmt.Printf("\tuser name:      %s\n", mqtt.user)
	fmt.Printf("\tpassword:  %s\n", mqtt.password)
	fmt.Printf("\ttopic to publish:     %s\n", mqtt.topic)
	fmt.Printf("\tqos:       %d\n", mqtt.qos)
	fmt.Printf("\tcleansess: %v\n", mqtt.cleansess)
	//fmt.Printf("\tnum:       %d\n", mqtt.num)
	fmt.Printf("\tstore:     %s\n", mqtt.store)

	//Passing the configured credentails to the client to start the publishing activity
	opts := MQTT.NewClientOptions()
	opts.AddBroker(mqtt.broker)
	opts.SetClientID(mqtt.id)
	opts.SetUsername(mqtt.user)
	opts.SetPassword(mqtt.password)
	opts.SetCleanSession(mqtt.cleansess)
	if mqtt.store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(mqtt.store))
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//loops up every one sec to gather the data and publishes them to the respected topic
loop:
	for {
		fmt.Println("*************************************Sample Publisher Started************************************************")

		fmt.Println("************************************doing publish*******************************************************")

		//loops on range of data ie result of generate function
		for payload := range generate() {

			//Publishing the gathered data to the topic
			client.Publish(mqtt.topic, byte(mqtt.qos), false, payload)

			lg.Infof("Published message %s", payload)
			lg.Infof("done...")
			continue loop
		}
	}

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")
}

func generate() <-chan string {
	c := make(chan string)
	go func() {

		var a [2]float32
		var retried int
		var err error

		a[0], a[1], retried, err = dht.ReadDHTxxWithRetry(dht.DHT11, 17, false, 10)
		timestamp := time.Now().Format("2019/00/00 00/00/00")

		if err != nil {
			fmt.Println(a[1])

		}
		//fmt.Println(a[0], a[1])
		lg.Infof("retried %d times", retried)

		c <- fmt.Sprintf(`{"temperature":{"value":%+v},"humidity":{"value":%+v},"timestamp":%+v}`, a[0], a[1], timestamp)

	}()

	return c
}
