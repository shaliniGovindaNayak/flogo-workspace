package redis

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/hoisie/redis"
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

	var client redis.Client
	var key = context.GetInput("keyval").(string)
	var value = context.GetInput("value").(string)
	client.Set(key, []byte(value))
	val, _ := client.Get(key)
	println(key, string(val))

	context.SetOutput("output", val)

	return true, nil
}
