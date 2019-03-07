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

	var tempValue float64

	sensorType := dht.DHT11

	//Fetch the data from the sensor dht11 from GPIO pin 17
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, 17, false, 10)
	if err != nil {
		//value = "false"
	}

	fmt.Println(temperature, humidity, retried)
	tempValue = float64(temperature)

	tupleDescAbsFileNm := common.GetAbsPathForResource("/home/isteer/flogo-workspace/rule/UnusualTemperatureBehaviour/TemperatureTupleDescription.json")
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

	//checking for temperature value
	rule := ruleapi.NewRule("Temperature is = 0")
	rule.AddCondition("c1", []string{"Temperature"}, checkForTempDec, nil)
	rule.SetAction(checkForTempDecAction)
	rule.SetContext("This is a test of context")
	rs.AddRule(rule)
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rule1 := ruleapi.NewRule("Temperature is > 100")
	rule1.AddCondition("c1", []string{"Temperature"}, checkForTempExc, nil)
	rule1.SetAction(checkForTempExcAction)
	rule1.SetContext("This is a test of context")
	rs.AddRule(rule1)
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rule2 := ruleapi.NewRule("Temperature is >0 and < 100")
	rule2.AddCondition("c1", []string{"Temperature"}, checkForTempNorm, nil)
	rule2.SetAction(checkForTempNormAction)
	rule2.SetContext("This is a test of context")
	rs.AddRule(rule2)
	fmt.Printf("Rule added: [%s]\n", rule.GetName())

	rs.Start(nil)

	//Now assert a "Temperature" tuple
	fmt.Println("Asserting Temperature tuple with Value=0")
	t1, _ := model.NewTupleWithKeyValues("Temperature", tempValue)
	t1.SetDouble(nil, "Value", tempValue)
	rs.Assert(nil, t1)

	//Retract tuples
	rs.Retract(nil, t1)

	//delete the rule
	rs.DeleteRule(rule.GetName())

	//unregister the session, i.e; cleanup
	rs.Unregister()
}

func checkForTempDec(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Temperature is 0 or less
	t1 := tuples["Temperature"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("Value")
	return name < 1
}

func checkForTempDecAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Temperature"]
	data, _ := t1.GetDouble("Value")

	if data < 1 {
		fmt.Println("====> Temperature is decreasing <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

func checkForTempExc(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Temperature is 100 or more
	t1 := tuples["Temperature"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("Value")
	return name >= 100
}

func checkForTempExcAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Temperature"]
	data, _ := t1.GetDouble("Value")

	if data >= 100 {
		fmt.Println("====> Temperature is exceeding <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

func checkForTempNorm(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	//This conditions filters on Temperature lies between 0 to 100
	t1 := tuples["Temperature"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name, _ := t1.GetDouble("Value")
	return name > 0 && name < 100
}

func checkForTempNormAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["Temperature"]
	data, _ := t1.GetDouble("Value")

	if data < 100 && data > 0 {
		fmt.Println("====> Temperature is Normal <=====")
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}
