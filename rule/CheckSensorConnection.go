package main

import (
	"context"
	"fmt"

	dht "github.com/d2r2/go-dht"
	logger "github.com/d2r2/go-logger"
	"github.com/fatih/color"
	"github.com/project-flogo/rules/common"
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

var lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
)

func main() {

	defer logger.FinalizeLogger()

	sensorType := dht.DHT11
	value := "true" //Set the value to be true by default

	//Fetch the data from the sensor dht11 from GPIO pin 17
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, 17, false, 10)
	if err != nil {
		value = "false"
	}

	fmt.Println(temperature, humidity, retried)

	//Load the tuple descriptor file (relative to GOPATH)
	tupleDescAbsFileNm := common.GetAbsPathForResource("SensorTupleDescriptor.json")
	tupleDescriptor := common.FileToString(tupleDescAbsFileNm)
	fmt.Printf("Loaded tuple descriptor: \n%s\n", tupleDescriptor)
	//First register the tuple descriptors
	err = model.RegisterTupleDescriptors(tupleDescriptor)
	if err != nil {
		fmt.Printf("Error [%s]\n", err)
		return
	}

	rs, _ := ruleapi.GetOrCreateRuleSession("asession") //Create a Rule session

	rule := ruleapi.NewRule("Sensor.Connection == true") //Configure the next rule,condition and the respective action to be performed
	rule.AddCondition("c1", []string{"Sensor"}, checkForConn, nil)
	rule.SetAction(checkForConnAction)
	rule.SetContext("This is a test of context")
	rs.AddRule(rule) //Adding the new rule to session
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rule = ruleapi.NewRule("Sensor.Connection == false") //Configure the next rule,condition and the respective action to be performed
	rule.AddCondition("c1", []string{"Sensor"}, checkForDisconn, nil)
	rule.SetAction(checkForDisconnAction)
	rule.SetContext("This is a test of context")
	rs.AddRule(rule) //Adding rule to session
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rs.Start(nil)

	fmt.Println("Asserting Sensor tuple with Connection=true") //Asserting the tuple value to "true", if true the rule is fired and Action is performed
	t1, _ := model.NewTupleWithKeyValues("Sensor", value)
	t1.SetString(nil, "Connection", value)
	rs.Assert(nil, t1)

	fmt.Println("Asserting Sensor tuple with Connection=false") //Asserting the tuple value to "false", if true the rule is fired and Action is performed
	t1, _ = model.NewTupleWithKeyValues("Sensor", value)
	t1.SetString(nil, "Connection", value)
	rs.Assert(nil, t1)

}

func checkForConn(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Connection = "true"

	t1 := tuples["Sensor"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	value, _ := t1.GetString("Connection")
	return value == "true"
}

func checkForConnAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Sensor"]

	//If Connection evalutes to be true then prints "sensor working fine"
	value, _ := t1.GetString("Connection")
	if value == "true" {
		fmt.Println()
		fmt.Println()
		fmt.Println(color.HiBlueString("###################################  Sensor is Connected and in Running State:):)   ####################################"))

		fmt.Println()
	}

	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

func checkForDisconn(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Connection = "false"
	t1 := tuples["Sensor"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	data, _ := t1.GetString("Connection")
	return data == "false"
}

func checkForDisconnAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Sensor"]

	////If Connection evalutes to be false then prints "sensor is disconnected"
	data, _ := t1.GetString("Connection")
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
