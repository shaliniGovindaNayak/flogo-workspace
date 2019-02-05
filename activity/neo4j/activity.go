package neo4j

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	gonorm "github.com/marpaia/GonormCypher"
)

var g *gonorm.Gonorm

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

	g = gonorm.New("http://192.168.1.34", 7474)
	result, err := g.Cypher(`
    CREATE (n:employee)
 `).Execute().AsString()

	if err != nil {
		panic(err)
	}

	fmt.Println("The result is:", result)
	context.SetOutput("output", result)
	return true, nil
}
