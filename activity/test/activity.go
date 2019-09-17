package test

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.Method)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

type Activity struct {
	settings *Settings
}

// Activity is an sample Activity that can be used as a base to create a custom activity

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	method := a.settings.Method
	fmt.Println(method)

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	num1 := input.Num1
	fmt.Println(num1)

	err = ctx.SetOutput("Output", input.Num2)
	if err != nil {
		return true, err
	}
	return true, nil
}
