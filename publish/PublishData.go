package main

import (
	"context"
	"fmt"

	dht "github.com/d2r2/go-dht"
	"github.com/fatih/color"
	"github.com/project-flogo/rules/common"
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"

	logger "github.com/d2r2/go-logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
)

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
	temp  string `json:"Temperature"`
	humid string `json:"Humidity"`
}

func main() {

}
func pub() {
	defer logger.FinalizeLogger()

	lg.Notify("***************************************************************************************************")
	lg.Notify("*** You can change verbosity of output, to modify logging level of module \"dht\"")
	lg.Notify("*** Uncomment/comment corresponding lines with call to ChangePackageLogLevel(...)")

	//MQTT credentails are assinged to variabl with respect to a specific host- 192.168.0.73(Raspberry pi)
	credentails := mqttParameters{"topic", "tcp://192.168.1.17:1883", "password", "username", "host", false, 0, 1, "pub", ":memory"}
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
		//for payload := range generate() {

		//Publishing the gathered data to the topic
		client.Publish(mqtt.topic, byte(mqtt.qos), false, payload)

		lg.Infof("Published message %s", payload)
		lg.Infof("done...")
		continue loop
		//}
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
		if err != nil {
			value := "false"
			rule(value)
		}
		//fmt.Println(a[0], a[1])
		lg.Infof("retried %d times", retried)

		c <- fmt.Sprintf(`{
			"temperature": {
				"PV": %v
			},
			"humidity": {
				"PV": %v
			}
		}`, a[0], a[1])

	}()

	return c
}

func rule(value string) {

	//Load the tuple descriptor file (relative to GOPATH)
	tupleDescAbsFileNm := common.GetAbsPathForResource("sensorconn.json") //Fetching the tuple structure
	tupleDescriptor := common.FileToString(tupleDescAbsFileNm)

	err := model.RegisterTupleDescriptors(tupleDescriptor) //Registers the tuple properties and displays error in case of failure
	if err != nil {
		fmt.Printf("Error [%s]\n", err)
		return
	}

	rs, _ := ruleapi.GetOrCreateRuleSession("asession") //Creates a rule session

	rule := ruleapi.NewRule("sensorConn.data == false")
	rule.AddCondition("c1", []string{"sensorConn"}, checkForDisconn, nil)
	rule.SetAction(checkForDisconnAction)
	rule.SetContext("This is a test of context")
	rs.AddRule(rule)
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rs.Start(nil) //starts the rule session

	fmt.Println("Asserting sensorConn tuple with data=false")
	t1, _ := model.NewTupleWithKeyValues("sensorConn", value)
	t1.SetString(nil, "data", value)
	rs.Assert(nil, t1)
}

func checkForDisconn(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on data="false"
	t1 := tuples["sensorConn"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	data, _ := t1.GetString("data")
	return data == "false"
}

func checkForDisconnAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	//This Action is triggered when the rule is fired

	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["sensorConn"]

	data, _ := t1.GetString("data")
	if data == "false" {
		fmt.Println()
		fmt.Println()
		fmt.Println(color.HiRedString("###################################  Sensor is DisConnected:(:(  ####################################"))
		fmt.Println()
		fmt.Println()
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}
