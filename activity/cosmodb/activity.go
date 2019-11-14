package cosmodb

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/project-flogo/core/activity"
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

type Activity struct {
}

// Activity is an sample Activity that can be used as a base to create a custom activity

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func insert(db string, pw string , un string , url string, data string) string{
	//username := "smartflo-iotdata"
	database := db
	password := pw

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{url}, // Get HOST + PORT
		//smartflo-iotdata:0E594yhEhx7UVptwtVGeAam5IOfLBcPMJzxFxDyo3TUjeOAI5wuPcTXRCgLomUnLhgo1KFcP1L5OQ7sDrsUvZA==@
		Timeout:  60 * time.Second,
		Database: database, // It can be anything
		Username: database, // Username
		Password: password, // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}else{
		fmt.Println("connected")
	}
	defer session.Close()

	session.SetSafe(&mgo.Safe{})
	collection := session.DB(database).C("details")

	// insert Document in collection
	// insert Document in collection
	err = collection.Insert(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}else{
		fmt.Println("inserted")
	}
	return "success"
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	fmt.Println(input.username)

	fmt.Println("requesting...")
	insert(input.database , input.username, input.password, input.url, input.data)
	//fmt.Println("insident raised")

	output := &Output{Output: " "}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}
	return true, nil
}
