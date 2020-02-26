package battery

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/distatus/battery"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Output{})

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

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func batteryDetails() string {
	batteries, err := battery.GetAll()
	//fmt.Println(batteries[0])
	if err != nil {
		fmt.Println("Could not get battery info!")

	}

	as, err := getMacAddr()
	jsondata := map[string]interface{}{
		"Battery":     batteries,
		"Time_stamp":  time.Now().UTC().Format("2006-01-02 15:04:05"),
		"Mac_address": as[3],
	}
	b, err := json.Marshal(jsondata)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	out := string(b)
	return out
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	data := batteryDetails()
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)

	if err != nil {
		log.Fatal(err)
	}
	//for _, a := range as {
	out := data
	fmt.Println(out)
	//log.Println("setting:", settings.ASetting)
	//ctx.Logger().Debug("Output: %s", settings.ASetting)
	//ctx.Logger().Debugf("Input: %s", input.AnInput)

	output := &Output{Output: out}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
