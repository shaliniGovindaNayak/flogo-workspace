package cosmodb

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	//utils "github.com/Azure/go-autorest/autorest"
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

type Details struct {
	Total_memory string
	Free_memory string
	Percentage_used_memory string
	Total_disk_space string
	Used_disk_space string
	Free_disk_space string
	Percentage_disk_space_usage string
	CPU_index_number string
	VendorID string
	Family string
	Speed string
	Uptime string
	Number_of_processes_running string
	Host_ID string
 }

func insertdata(username string, url string, password string, content map[string]interface{}){

	database := username
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{url}, // Get HOST + PORT
		//smartflo-iotdata:0E594yhEhx7UVptwtVGeAam5IOfLBcPMJzxFxDyo3TUjeOAI5wuPcTXRCgLomUnLhgo1KFcP1L5OQ7sDrsUvZA==@
		Timeout:  1000 * time.Second,
		Database: database, // It can be anything
		Username: database, // Username
		Password: password, // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	fmt.Println(dialInfo.Timeout)
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

	//cont := content.(Details)

	// insert Document in collection
	// insert Document in collection
	err = collection.Insert(content)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}else{
		fmt.Println("inserted")
	}
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	fmt.Println(input.Username)

	fmt.Println("requesting...")
	insertdata(input.Username, input.Connectionstring, input.Password, input.Content)
	output := &Output{Output: "success"}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}
	return true, nil
}
