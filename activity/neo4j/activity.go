package neo4j

import (
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
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

	url := context.GetInput("url").(string)
	driver := bolt.NewDriver()
	log.Println("created new driver")
	conn, err := driver.Open(url)
	if err != nil {
		log.Println("error while trying to connect to bolt")
	}
	result, err := conn.Prepare("create (n:node) return n")
	if err != nil {
		log.Println("error while executing the query")
	}

	log.Println("output:", result)
	context.SetOutput("output", result)

	return true, nil
}
