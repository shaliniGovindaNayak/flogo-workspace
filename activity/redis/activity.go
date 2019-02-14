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

	key := context.GetInput("key").(string)
	var client redis.Client
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush(key, []byte(v))
	}
	var out []string
	dbvals, _ := client.Lrange(key, 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
		out[i] = string(v)
	}

	context.SetOutput("output", out)

	return true, nil
}
