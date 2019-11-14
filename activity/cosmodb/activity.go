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

func insert(db string, pw string , un string , url string) string{
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

	// insert Document in collection
	// insert Document in collection
	err = collection.Insert(&Details{Total_memory:"250966470656",Free_memory:"311087104",Percentage_used_memory:"62.57",Total_disk_space:"250966470656",Used_disk_space:"43844227072",Free_disk_space:"194302513152",Percentage_disk_space_usage:"18.41",CPU_index_number:"0",VendorID:"GenuineIntel",Family:"6",Speed:"2900.00",Uptime:"235282",Number_of_processes_running:"354",Host_ID:"74619e31-ba1c-45c9-9473-c4cc05c0b558"})
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
	fmt.Println(input.Username)

	fmt.Println("requesting...")
	out := insert(input.Database , input.Username, input.Password, input.Url)
	//fmt.Println("insident raised")

	output := &Output{Output: out }

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}
	return true, nil
}
