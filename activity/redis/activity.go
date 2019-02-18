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
	var res []string

	switch operation {

	case "strings":
		result = set(key, value)
		context.SetOutput("output", result)
		break

	case "hash":
		result = hash(key, field, value)
		context.SetOutput("output", result)
		break

	case "list":
		val := []string{"hai", "hello", "bye"}
		res = list(key, val)
		context.SetOutput("output", res)
		break

	}

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

func list(key string, value []string) []string {

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	var res []string
	for i := 0; i <= len(value); i++ {
		redis.Lpush(key, value[i])
	}

	res, _ = redis.List(key)
	return res
}
