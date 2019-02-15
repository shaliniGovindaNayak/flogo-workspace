package redisactivity

import (
	"encoding/json"

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

var input struct {
	name   string
	age    string
	salary string
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	key := context.GetInput("key").(string)
	operation := context.GetInput("operation").(string)
	data := context.GetInput("data").(string)

	var field []string
	field[0] = "name"
	field[1] = "age"
	field[2] = "salary"

	var value []string

	in, _ := json.Marshal(data)
	json.Unmarshal(in, &input)

	value[0] = input.name
	value[1] = input.age
	value[2] = input.salary

	var result string
	switch operation {

	case "strings":
		result = set(key, value[0])
		break

	case "hash":
		for i := 0; i < len(field); i++ {
			result = hash(key, field[i], value[i])

		}
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
