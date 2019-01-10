package stringtojson

import (
	"encoding/json"
	"fmt"

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

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	type Data struct {
		Temp  string `json:"Temp"`
		Humid string `json:"Humid"`
	}

	input := context.GetInput("Rawstring").(string)
	println(input)
	in := []byte(input)

	u1 := Data{}
	if err := json.Unmarshal(in, &u1); err != nil {
		//log.Fatal(err)
	}
	fmt.Println("Temperature:", u1.Temp)
	fmt.Println("Humidity:", u1.Humid)
	//out := u1
	context.SetOutput("Json[0]", u1.Temp)
	context.SetOutput("Json[1]", u1.Humid)
	return true, nil
}
