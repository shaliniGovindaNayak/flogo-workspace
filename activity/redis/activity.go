package redis

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
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

	Commands := context.GetInput("Commands").(string)

	key := context.GetInput("key").(string)
	value := context.GetInput("value").(string)
	incr := context.GetInput("Incr").(int)
	//decr := context.GetInput("Decr").(int)
	field := context.GetInput("field").(string)

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	switch Commands {

	case "TTL":
		res := string(redis.TTL(key))
		context.SetOutput("output", res)
		break

	//case "SetTTL":
	//res = setTTl(key)

	case "Type":
		res := redis.Type(key)
		logger.Debug(res)
		context.SetOutput("output", res)
		break

	case "Dump":
		res := redis.Dump()
		context.SetOutput("output", res)
		break

	case "Del":
		res := redis.Del(key)
		context.SetOutput("output", res)
		break

	//case "DB":
	//	res = db()
	//	break

	case "Exists":
		res := redis.Exists(key)
		context.SetOutput("output", res)
		break

	case "Flushall":
		redis.FlushAll()
		break

	case "FlushDB":
		redis.FlushDB()
		break

	case "Set":
		redis.Set(key, value)
		context.SetOutput("output", "done")
		break

	case "Get":
		res, _ := redis.Get(key)
		context.SetOutput("output", res)

	case "Increment":
		res, _ := redis.Incr(key, incr)
		context.SetOutput("output", res)

	case "Hget":
		res := redis.HGet(key, field)
		context.SetOutput("output", res)
		break

	case "Hset":
		redis.HSet(key, field, value)
		break

	case "Hdel":
		redis.HDel(key, field)
		break

		//	case "Hkeys":
	//	res := redis.HKeys(key)
	//		break

	case "Hincrement":
		res, _ := redis.HIncr(key, field, incr)
		context.SetOutput("output", res)

		//	case "List":
		//		res = redis.List(key)
		//		break

	case "Lpush":
		res, _ := redis.Lpush(key, value)
		context.SetOutput("output", res)
		break

	case "Lpop":
		res, _ := redis.Lpop(key)
		context.SetOutput("output", res)
		break

	case "Rpush":
		res, _ := redis.Push(key, value)
		context.SetOutput("output", res)
		break

	case "Rpop":
		res, _ := redis.Pop(key)
		context.SetOutput("output", res)
		break

	case "Add":
		res, _ := redis.SetAdd(key, value)
		context.SetOutput("output", res)
		break

		//	case "Members":
		//		res = redis.Members(key)

		//	case "Zadd":
		//		res = redis.ZAdd(key)

	case "Zrem":
		res, _ := redis.ZRem(key, value)
		context.SetOutput("output", res)

	case "Zscore":
		res, _ := redis.ZScore(key, value)
		context.SetOutput("output", res)

	case "Ismember":
		res, _ := redis.IsMember(key, value)
		context.SetOutput("output", res)
	}

	return true, nil
}
