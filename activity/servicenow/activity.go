package servicenow

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.Username)
	ctx.Logger().Debugf("Setting: %s", s.Password)
	ctx.Logger().Debugf("Setting: %s", s.Instanceurl)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

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
	Settings := &Settings{}
	Output := &Output{}
	ctx.GetInputObject(input)

	username := Settings.Username
	password := Settings.Password
	instanceURL := Settings.Instanceurl
	insidentVALUE := input.Content

	fmt.Println(insidentVALUE)

	fmt.Println("requesting...")
	S := basicAuth(username, password, instanceURL, insidentVALUE)
	fmt.Println(S)
	fmt.Println("insident raised")
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	Output.Output = "sucess"

	return true, nil
}
