package stringtojson

import (
	"encoding/json"
	"fmt"

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

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	input := context.GetInput("Rawstring").(string)
	println(input)

	type Data struct {
		Temp  string
		Humid string
	}
	var datas Data
	err = json.Unmarshal([]byte(input), &datas)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", datas)
	context.SetOutput("Json", datas)

	return true, nil
}
