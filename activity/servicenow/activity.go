package servicenow

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance
	return act, nil
}

type Activity struct {
}

// Activity is an sample Activity that can be used as a base to create a custom activity

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func basicAuth(username string, password string, instanceURL string, instanceVALUE string) string {

	client := &http.Client{}

	req, err := http.NewRequest("POST", instanceURL, bytes.NewBufferString(instanceVALUE))
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	username := input.Username
	password := input.Password
	instanceurl := input.Instanceurl

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	incidentvalue := input.Content
	fmt.Println(input.Username)

	fmt.Println("requesting...")
	basicAuth(username, password, instanceurl, incidentvalue)
	fmt.Println("insident raised")

	output := &Output{Output: input.Content}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}
	return true, nil
}
