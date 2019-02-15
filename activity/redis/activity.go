package redis

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/alicebob/miniredis"
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
	value := context.GetInput("value").(string)
	operation := context.GetInput("operation").(string)
	field := context.GetInput("field").(string)
	var result string
	switch operation {

	case "strings":
		result = set(key, value)
		break

	case "hash":
		result = hash(key, field, value)
		break

	case "list":
		result = string(list(key, value))
		break

	}
	context.SetOutput("output", result)

	return true, nil
}

func set(key string, value string) string {

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	redis.Set(key, value)
	res, _ := redis.Get(key)

	return res

}

func hash(key string, field string, value string) string {

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	redis.HSet(key, field, value)
	return redis.HGet(key, field)

}

func list(key string, value string) int {

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	res, _ := redis.Push(key, value)
	return res
}
