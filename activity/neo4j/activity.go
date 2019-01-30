package neo4j

import (
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"

	neo4j "github.com/davemeehan/Neo4j-GO"
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

	// do eval

	node := map[string]string{
		"test1": "foo",
		"test2": "bar",
	}

	n, err := neo4j.NewNeo4j("http://localhost:7474/user/neo4j", "neo4j", "password")

	data, _ := n.CreateNode(node)
	log.Printf("\nNode ID: %v\n", data.ID)
	self := data.ID

	data, _ = n.GetNode(self)
	log.Printf("\nNode data: %v\n", data)

	context.SetOutput("output", data.ID)

	return true, nil
}
