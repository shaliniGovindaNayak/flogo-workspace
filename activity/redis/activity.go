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

	Commands := context.GetInput("Commands").(string)

	key := context.GetInput("key").(string)
	value := context.GetInput("value").(string)
	incr := context.GetInput("Incr").(int)
	//decr := context.GetInput("Decr").(int)
	field := context.GetInput("field").(string)

	var result1 string
	var result2 int
	var result3 bool

	redis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	switch Commands {

	case "TTL":
		result1 = string(redis.TTL(key))
		break

	//case "SetTTL":
	//res = setTTl(key)

	case "Type":
		result1 = redis.Type(key)
		break

	case "Dump":
		result1 = redis.Dump()
		break

	case "Del":
		result3 = redis.Del(key)
		break

	//case "DB":
	//	res = db()
	//	break

	case "Exists":
		result3 = redis.Exists(key)
		break

	case "Flushall":
		redis.FlushAll()
		result1 = "done"
		break

	case "FlushDB":
		redis.FlushDB()
		result1 = "done"
		break

	case "Set":
		redis.Set(key, value)
		result1 = "done"
		break

	case "Get":
		result1, _ = redis.Get(key)
		break

	case "Increment":
		result2, _ = redis.Incr(key, incr)
		break

	case "Hget":
		result1 = redis.HGet(key, field)
		break

	case "Hset":
		redis.HSet(key, field, value)
		result1 = "done"
		break

	case "Hdel":
		redis.HDel(key, field)
		result1 = "done"
		break

		//	case "Hkeys":
	//	res := redis.HKeys(key)
	//		break

	//case "Hincrement":
	//result3, _ = redis.HIncr(key, field, incr)

	//	case "List":
	//		res = redis.List(key)
	//		break

	case "Lpush":
		result2, _ = redis.Lpush(key, value)
		break

	case "Lpop":
		result1, _ = redis.Lpop(key)
		break

	case "Rpush":
		result2, _ = redis.Push(key, value)
		break

	case "Rpop":
		result1, _ = redis.Pop(key)
		break

	case "Add":
		result2, _ = redis.SetAdd(key, value)
		break

		//	case "Members":
		//		res = redis.Members(key)

		//	case "Zadd":
		//		res = redis.ZAdd(key)

	case "Zrem":
		result3, _ = redis.ZRem(key, value)
		break

	//case "Zscore":
	//	res, _ := redis.ZScore(key, value)
	//	break

	case "Ismember":
		result3, _ = redis.IsMember(key, value)
	}

	context.SetOutput("output1", result1)
	context.SetOutput("output2", result2)
	context.SetOutput("output", result3)

	return true, nil
}
