package servicenow

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

type Input struct {
	Username      string `md:"Username"`    // The message to log
	Password      string `md:"Password"`    // Append contextual execution information to the log message
	Instanceurl   string `md:"Instanceurl"` // The message to log
	insidentvalue string `md:"insidentvalue"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Username":      i.Username,
		"Password":      i.Password,
		"Instanceurl":   i.Instanceurl,
		"insidentvalue": i.insidentvalue,
	}
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Username, err = coerce.ToString(values["Username"])
	if err != nil {
		return err
	}
	i.Password, err = coerce.ToString(values["Password"])
	if err != nil {
		return err
	}
	i.Instanceurl, err = coerce.ToString(values["Instanceurl"])
	if err != nil {
		return err
	}
	i.insidentvalue, err = coerce.ToString(values["insidentvalue"])
	if err != nil {
		return err
	}

	return nil
}

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

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
	ctx.GetInputObject(input)

	username := input.Username
	password := input.Password
	instanceURL := input.Instanceurl
	insidentVALUE := input.insidentvalue

	fmt.Println("requesting...")
	S := basicAuth(username, password, instanceURL, insidentVALUE)
	fmt.Println(S)
	fmt.Println("insident raised")
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	output := "sucess"

	return true, nil
}
