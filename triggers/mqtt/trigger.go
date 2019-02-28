package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/rules/common"
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-mqtt")

// MqttTrigger is simple MQTT trigger
type MqttTrigger struct {
	metadata       *trigger.Metadata
	client         mqtt.Client
	config         *trigger.Config
	handlers       []*trigger.Handler
	topicToHandler map[string]*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MQTTFactory{metadata: md}
}

// MQTTFactory MQTT Trigger factory
type MQTTFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *MQTTFactory) New(config *trigger.Config) trigger.Trigger {
	return &MqttTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *MqttTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize implements trigger.Initializable.Initialize
func (t *MqttTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

// Start implements trigger.Trigger.Start
func (t *MqttTrigger) Start() error {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(t.config.GetSetting("broker"))
	opts.SetClientID(t.config.GetSetting("id"))
	opts.SetUsername(t.config.GetSetting("user"))
	opts.SetPassword(t.config.GetSetting("password"))
	b, err := data.CoerceToBoolean(t.config.Settings["cleansess"])
	if err != nil {
		log.Error("Error converting \"cleansess\" to a boolean ", err.Error())
		return err
	}
	opts.SetCleanSession(b)
	if storeType := t.config.Settings["store"]; storeType != ":memory:" {
		opts.SetStore(mqtt.NewFileStore(t.config.GetSetting("store")))
	}

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		//TODO we should handle other types, since mqtt message format are data-agnostic
		payload := string(msg.Payload())
		log.Debug("Received msg:", payload)
		handler, found := t.topicToHandler[topic]
		if found {
			t.RunHandler(handler, payload)
		} else {
			log.Errorf("handler for topic '%s' not found", topic)
		}
	})

	client := mqtt.NewClient(opts)
	t.client = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	i, err := data.CoerceToDouble(t.config.Settings["qos"])
	if err != nil {
		log.Error("Error converting \"qos\" to an integer ", err.Error())
		return err
	}

	t.topicToHandler = make(map[string]*trigger.Handler)

	for _, handler := range t.handlers {

		topic := handler.GetStringSetting("topic")

		if token := t.client.Subscribe(topic, byte(i), nil); token.Wait() && token.Error() != nil {
			log.Errorf("Error subscribing to topic %s: %s", topic, token.Error())
			return token.Error()
		} else {
			log.Debugf("Subscribed to topic: %s, will trigger handler: %s", topic, handler)
			t.topicToHandler[topic] = handler
		}

		fmt.Println("** rulesapp: Example usage of the Rules module/API **")

		//Load the tuple descriptor file (relative to GOPATH)
		tupleDescAbsFileNm := common.GetAbsPathForResource("https://github.com/shaliniGovindaNayak/flogo-workspace/tree/master/triggers/mqtt/example.json")
		tupleDescriptor := common.FileToString(tupleDescAbsFileNm)
		fmt.Printf("Loaded tuple descriptor: \n%s\n", tupleDescriptor)
		//First register the tuple descriptors
		err = model.RegisterTupleDescriptors(tupleDescriptor)
		if err != nil {
			fmt.Printf("Error [%s]\n", err)

		}

		rs, _ := ruleapi.GetOrCreateRuleSession("asession")
		//s, _ := ruleapi.GetOrCreateRuleSession("asession")

		rule := ruleapi.NewRule("MqttSubscribe.data==null")

		rule.AddCondition("c1", []string{"MqttSubscribe"}, checkForMqttSubscribeData, nil)
		rule.SetAction(checkForMqttSubscribeDataAction)
		rule.SetContext("This is a test of context")
		rs.AddRule(rule)
		fmt.Printf("Rule added: [%s]\n", rule.GetName())

		rs.Start(nil)
		fmt.Println("Asserting MqttSubscribe tuple with value=null")
		t1, _ := model.NewTupleWithKeyValues("MqttSubscribe", topic)
		t1.SetString(nil, "data", topic)
		rs.Assert(nil, t1)

		rs.Retract(nil, t1)
	}
	return nil
}

func checkForMqttSubscribeData(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {

	t1 := tuples["MqttSubscribe"]
	if t1 == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	data, _ := t1.GetString("data")

	return data == ""
}

func checkForMqttSubscribeDataAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	fmt.Printf("Context is [%s]\n", ruleCtx)
	t1 := tuples["MqttSubscribe"]
	data, _ := t1.GetString("data")
	if data == "" {
		fmt.Println("data is empty")
	} else {
		fmt.Println("data:%s", data)
	}
	if t1 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return
	}
}

// Stop implements ext.Trigger.Stop
func (t *MqttTrigger) Stop() error {
	//unsubscribe from topic
	for _, handlerCfg := range t.config.Handlers {
		log.Debug("Unsubscribing from topic: ", handlerCfg.GetSetting("topic"))
		if token := t.client.Unsubscribe(handlerCfg.GetSetting("topic")); token.Wait() && token.Error() != nil {
			log.Errorf("Error unsubscribing from topic %s: %s", handlerCfg.Settings["topic"], token.Error())
		}
	}

	t.client.Disconnect(250)

	return nil
}

// RunHandler runs the handler and associated action
func (t *MqttTrigger) RunHandler(handler *trigger.Handler, payload string) {

	trgData := make(map[string]interface{})
	trgData["message"] = payload

	results, err := handler.Handle(context.Background(), trgData)

	if err != nil {
		log.Error("Error starting action: ", err.Error())
	}

	log.Debugf("Ran Handler: [%s]", handler)

	var replyData interface{}

	if len(results) != 0 {
		dataAttr, ok := results["data"]
		if ok {
			replyData = dataAttr.Value()
		}
	}

	if replyData != nil {
		dataJson, err := json.Marshal(replyData)
		if err != nil {
			log.Error(err)
		} else {
			replyTo := handler.GetStringSetting("topic")
			if replyTo != "" {
				t.publishMessage(replyTo, string(dataJson))
			}
		}
	}
}

func (t *MqttTrigger) publishMessage(topic string, message string) {

	log.Debug("ReplyTo topic: ", topic)
	log.Debug("Publishing message: ", message)

	qos, err := strconv.Atoi(t.config.GetSetting("qos"))
	if err != nil {
		log.Error("Error converting \"qos\" to an integer ", err.Error())
		return
	}
	if len(topic) == 0 {
		log.Warn("Invalid empty topic to publish to")
		return
	}
	token := t.client.Publish(topic, byte(qos), false, message)
	sent := token.WaitTimeout(5000 * time.Millisecond)
	if !sent {
		// Timeout occurred
		log.Errorf("Timeout occurred while trying to publish to topic '%s'", topic)
		return
	}
}
