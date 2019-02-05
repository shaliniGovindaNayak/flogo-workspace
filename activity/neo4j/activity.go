package neo4j

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	gonorm "github.com/marpaia/GonormCypher"
)

//var url string
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

	operation := context.GetInput("operation").(string)
	query := context.GetInput("query").(string)

	switch operation {

	case "create":
		result, err := g.Cypher("`" + query + "`").Execute().AsString()
		if err != nil {
			panic(err)
		}
		context.SetOutput("output", result)

	case "read":
		result, err := g.Cypher("`" + query + "`").Execute().AsString()
		if err != nil {
			panic(err)
		}
		context.SetOutput("output", result)

	case "delete":
		result, err := g.Cypher("`" + query + "`").Execute().AsString()
		if err != nil {
			panic(err)
		}
		context.SetOutput("output", result)

	default:
		logger.Debug("invalid operation")

	}

	return true, nil
}

func init() {

	g = gonorm.New("http://neo4j:password@192.168.1.34", 7474)
}
