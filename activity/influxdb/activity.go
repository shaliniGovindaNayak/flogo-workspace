package influxdb

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client"         // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2" // this is important because of the bug in go mod
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

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	fmt.Println(input.Host)
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: input.Host,
	})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	var res []client.Result
	query := "SELECT * FROM " + input.Table
	q := client.NewQuery(query, input.Schema, "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
		res = response.Results
		fmt.Println(res)
	}

	output := &Output{Output: "sucess"}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
