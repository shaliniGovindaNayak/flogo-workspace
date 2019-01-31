package neo4j

import (
	"fmt"

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
	query := context.GetInput("query statement").(string)

	driver := bolt.NewDriver()
	conn, _ := driver.OpenNeo(url)
	defer conn.Close()

	// Start by creating a node
	result, _ := conn.ExecNeo(query, map[string]interface{}{"foo": 1, "bar": 2.2})
	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1
	context.SetOutput("result", numResult)

	return true, nil
}
