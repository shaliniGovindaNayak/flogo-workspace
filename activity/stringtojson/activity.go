package stringtojson

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-string2json")

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

type Data struct {
	temperature []data `json:temperature`
	humidity    []data `json:humidity`
}

type data struct {
	HV int `json:HV`
	LV int `json.LV`
	PV int `json.PV`
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	input := context.GetInput("Rawstring").(string)
	println(input)
	in := []byte(input)
	println(in)
	raw := make(map[string]interface{})
	json.Unmarshal(in, &raw)
	log.Debugf("the raw string")
	raw["count"] = 1
	out, _ := json.Marshal(&raw)

	log.Infof("the output value ... %s", string(out))
	context.SetOutput("Json", string(out))

	return true, nil
}
