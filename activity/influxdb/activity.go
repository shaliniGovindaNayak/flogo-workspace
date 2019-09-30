package influxdb

import (
	
	"time"
	"fmt"
	"github.com/project-flogo/core/activity"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
   client "github.com/influxdata/influxdb1-client/v2"
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

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: input.Host,
	})

	json := input.Values

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  input.Schema,
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{}
	fields := json
	pt, err := client.NewPoint(input.Table, tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Logger().Debugf("Input: %s", input.Input)

	output := &Output{Output: "ok"}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
