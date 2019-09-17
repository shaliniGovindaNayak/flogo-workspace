package servicenowv2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-serviceNowV2")

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
func basicAuth(username string, password string, instanceURL string, instanceVALUE string) string {

	client := &http.Client{}

	req, err := http.NewRequest("POST", instanceURL, bytes.NewBufferString(instanceVALUE))
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	username := context.GetInput("username").(string)
	password := context.GetInput("password").(string)
	instanceurl := context.GetInput("instanceurl").(string)
	incidentvalue := context.GetInput("content").(string)
	fmt.Println(username)

	fmt.Println("requesting...")
	S := basicAuth(username, password, instanceurl, incidentvalue)
	//fmt.Println(S)
	fmt.Println("insident raised")

	context.SetOutput("output", S)
	return true, nil

}
