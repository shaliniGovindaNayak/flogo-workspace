package main

import (
	"context"
	"fmt"

	dht "github.com/d2r2/go-dht"
	"github.com/project-flogo/rules/common"
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

func main() {
	fmt.Println("** rulesapp: Example usage of the Rules module/API **")

	var humidValue float64

	sensorType := dht.DHT11

	//Fetch the data from the sensor dht11 from GPIO pin 17
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, 17, false, 10)
	if err != nil {
		//value = "false"
	}

	fmt.Println(temperature, humidity, retried)
	humidValue = float64(humidity)

	tupleDescAbsFileNm := common.GetAbsPathForResource("/home/isteer/flogo-workspace/rule/UnusualHumidityBehaviour/HumidityTupleDescriptor.json")
	tupleDescriptor := common.FileToString(tupleDescAbsFileNm)

	fmt.Printf("Loaded tuple descriptor: \n%s\n", tupleDescriptor)
	//First register the tuple descriptors
	err = model.RegisterTupleDescriptors(tupleDescriptor)
	if err != nil {
		fmt.Printf("Error [%s]\n", err)
		return
	}

	//Create a RuleSession
	rs, _ := ruleapi.GetOrCreateRuleSession("asession")

	//checking for humidity value
	rule := ruleapi.NewRule("Humidity is = 0")
	rule.AddCondition("c1", []string{"Humidity"}, checkForHumidDec, nil)
	rule.SetAction(checkForHumidDecAction)
	rule.SetContext("This is a test of context")
	rs.AddRule(rule)
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rule1 := ruleapi.NewRule("Humidity is > 100")
	rule1.AddCondition("c1", []string{"Humidity"}, checkForHumidExc, nil)
	rule1.SetAction(checkForHumidExcAction)
	rule1.SetContext("This is a test of context")
	rs.AddRule(rule1)
	fmt.Printf("Rule added: [%s]\n", rule1.GetName())

	rule2 := ruleapi.NewRule("Humidity is >0 and < 100")
	rule2.AddCondition("c1", []string{"Humidity"}, checkForHumidNorm, nil)
	rule2.SetAction(checkForHumidNormAction)
	rule2.SetContext("This is a test of context")
	rs.AddRule(rule2)
	fmt.Printf("Rule added: [%s]\n", rule2.GetName())

	//Start the rule session
	rs.Start(nil)

	//	fmt.Println("Asserting Humidity tuple with Value=0")
	t2, _ := model.NewTupleWithKeyValues("Humidity", humidValue)
	t2.SetDouble(nil, "value", humidValue)
	rs.Assert(nil, t2)

	//delete the rule

	//unregister the session, i.e; cleanup
	rs.Unregister()
}

func checkForHumidDec(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Humidity is 0 or less
	t1 := tuples["Humidity"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("value")
	return name < 1
}

func checkForHumidDecAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Humidity"]
	data, _ := t1.GetDouble("value")

	if data < 1 {
		fmt.Println("====> Humidity is decreasing <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

func checkForHumidExc(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on humidity is 100 or more
	t1 := tuples["Humidity"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("value")
	return name >= 100
}

func checkForHumidExcAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Humidity"]
	data, _ := t1.GetDouble("value")

	if data >= 100 {
		fmt.Println("====> Humidity is exceeding <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

func checkForHumidNorm(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Humidity lies between 0 to 100
	t1 := tuples["Humidity"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("value")
	return name > 0 && name < 100
}

func checkForHumidNormAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Humidity"]
	data, _ := t1.GetDouble("value")

	if data < 100 && data >= 1 {
		fmt.Println("====> Humidity is Normal <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}
