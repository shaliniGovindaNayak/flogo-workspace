package neo4j

import (
	"fmt"
	"io"
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

const (
	uRI          = "bolt://neo4j:password@localhost:7687"
	createNode   = "CREATE (n:NODE {foo: {foo}, bar: {bar}})"
	getNode      = "MATCH (n:NODE) RETURN n.foo, n.bar"
	relationNode = "MATCH path=(n:NODE)-[:REL]->(m) RETURN path"
	deleteNodes  = "MATCH (n) DETACH DELETE n"
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

	fmt.Println(url)

	con := createConnection()
	defer con.Close()

	st := prepareSatement(createNode, con)
	executeStatement(st)

	st = prepareSatement(getNode, con)
	rows := queryStatement(st)
	consumeRows(rows, st)

	pipe := preparePipeline(con)
	executePipeline(pipe)

	st = prepareSatement(relationNode, con)
	rows = queryStatement(st)
	consumeMetadata(rows, st)

	cleanUp(deleteNodes, con)

	return true, nil
}

func createConnection() bolt.Conn {
	driver := bolt.NewDriver()
	con, err := driver.OpenNeo(uRI)
	handleError(err)
	return con
}

// Here we prepare a new statement. This gives us the flexibility to
// cancel that statement without any request sent to Neo
func prepareSatement(query string, con bolt.Conn) bolt.Stmt {
	st, err := con.PrepareNeo(query)
	handleError(err)
	return st
}

// Here we prepare a new pipeline statement for running multiple
// queries concurrently
func preparePipeline(con bolt.Conn) bolt.PipelineStmt {
	pipeline, err := con.PreparePipeline(
		"MATCH (n:NODE) CREATE (n)-[:REL]->(f:FOO)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(b:BAR)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(z:BAZ)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(f:FOO)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(b:BAR)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(z:BAZ)",
	)
	handleError(err)
	return pipeline
}

func executePipeline(pipeline bolt.PipelineStmt) {
	pipelineResults, err := pipeline.ExecPipeline(nil, nil, nil, nil, nil, nil)
	handleError(err)

	for _, result := range pipelineResults {
		numResult, _ := result.RowsAffected()
		fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 2 (per each iteration)
	}

	err = pipeline.Close()
	handleError(err)
}

func queryStatement(st bolt.Stmt) bolt.Rows {
	// Even once I get the rows, if I do not consume them and close the
	// rows, Neo will discard and not send the data
	rows, err := st.QueryNeo(nil)
	handleError(err)
	return rows
}

func consumeMetadata(rows bolt.Rows, st bolt.Stmt) {
	// Here we loop through the rows until we get the metadata object
	// back, meaning the row stream has been fully consumed

	var err error
	err = nil

	for err == nil {
		var row []interface{}
		row, _, err = rows.NextNeo()
		if err != nil && err != io.EOF {
			panic(err)
		} else if err != io.EOF {
			fmt.Printf("PATH: %#v\n", row[0].(graph.Path)) // Prints all paths
		}
	}
	st.Close()
}

func consumeRows(rows bolt.Rows, st bolt.Stmt) {
	// This interface allows you to consume rows one-by-one, as they
	// come off the bolt stream. This is more efficient especially
	// if you're only looking for a particular row/set of rows, as
	// you don't need to load up the entire dataset into memory
	data, _, err := rows.NextNeo()
	handleError(err)

	// This query only returns 1 row, so once it's done, it will return
	// the metadata associated with the query completion, along with
	// io.EOF as the error
	_, _, err = rows.NextNeo()
	handleError(err)
	fmt.Printf("COLUMNS: %#v\n", rows.Metadata()["fields"].([]interface{})) // COLUMNS: n.foo,n.bar
	fmt.Printf("FIELDS: %d %f\n", data[0].(int64), data[1].(float64))       // FIELDS: 1 2.2

	st.Close()
}

// Executing a statement just returns summary information
func executeStatement(st bolt.Stmt) {
	result, err := st.ExecNeo(map[string]interface{}{"foo": 1, "bar": 2.2})
	handleError(err)
	numResult, err := result.RowsAffected()
	handleError(err)
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	// Closing the statment will also close the rows
	st.Close()
}

func cleanUp(query string, con bolt.Conn) {
	result, _ := con.ExecNeo(query, nil)
	fmt.Println(result)
	numResult, _ := result.RowsAffected()
	fmt.Printf("Rows Deleted: %d", numResult) // Rows Deleted: 13
}

// Here we create a simple function that will take care of errors, helping with some code clean up
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
