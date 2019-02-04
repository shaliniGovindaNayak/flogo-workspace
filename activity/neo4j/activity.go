package neo4j

import (
	"database/sql"
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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

	log.Println("activity starts")
	url := context.GetInput("url").(string)

	log.Fatal("fetching input")

	db, err := sql.Open("neo4j - cyper", url)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`create (n:employee)`)
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("output:", rows)
	context.SetOutput("output", rows)

	return true, nil
}
