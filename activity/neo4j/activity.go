package neo4j

import (
	"database/sql"

	"log"

	_ "gopkg.in/cq.v1"

	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var logg = logger.GetLogger("activity-neo4j")

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
	db, err := sql.Open("neo4j-cypher", url)
	logg.Debug(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("create (:User {screenName:{0}})")
	if err != nil {
		log.Fatal(err)
	}
	logg.Debug(stmt)

	stmt.Exec("wefreema")
	stmt.Exec("JnBrymn")
	stmt.Exec("technige")

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	logg.Debug(err)

	return true, nil
}
