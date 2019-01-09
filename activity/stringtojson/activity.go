package stringtojson

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

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

var raw = make(map[string]interface{})

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	input := context.GetInput("Rawstring").(string)
	println(input)
	in := []byte(input)

	json.Unmarshal(in, &raw)
	raw["count"] = 1
	out, _ := json.Marshal(raw)
	println(string(out))
	context.SetOutput("Json", out)

	return true, nil
}
