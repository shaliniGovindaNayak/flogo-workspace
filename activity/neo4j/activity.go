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

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://localhost:7687")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareNeo("CREATE (n:NODE {foo: {foo}, bar: {bar}})")
	if err != nil {
		panic(err)
	}

	// Executing a statement just returns summary information
	result, err := stmt.ExecNeo(map[string]interface{}{"foo": 1, "bar": 2.2})
	if err != nil {
		panic(err)
	}
	numResult, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	context.SetOutput("output", numResult)
	// Closing the statment will also close the rows
	stmt.Close()

	return true, nil
}
