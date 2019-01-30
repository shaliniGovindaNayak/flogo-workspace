package neo4j

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	n "github.com/davemeehan/Neo4j-GO"
	//n "github.com/go-cq/cq"
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

	user := context.GetInput("username").(string)
	pass := context.GetInput("password").(string)
	resp, err := n.NewNeo4j("", user, pass)

	context.SetOutput("resp", resp)

	return true, nil
}
