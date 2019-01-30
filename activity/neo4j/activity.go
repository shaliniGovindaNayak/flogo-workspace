package neo4j

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

//type meth string
type neo4j struct {
	Method     string // which http method
	StatusCode int    // last http status code received
	URL        string
	Username   string
	Password   string
}

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

	meth := context.GetInput("method").(string)
	user := context.GetInput("username").(string)
	pass := context.GetInput("password").(string)
	resp, err := newNeo4j("", user, pass, meth)
	context.SetOutput("resp", resp)
	return true, nil
}

func newNeo4j(u string, user string, passwd string, method string) (*neo4j, error) {
	n := new(neo4j)
	if len(u) < 1 {
		u = "http://192.168.1.34:7474/user/neo4j"
	}
	if len(user) > 0 {
		n.Username = user
	}
	if len(passwd) > 0 {
		n.Password = passwd
	}

	n.URL = u
	_, err := n.send(u, "", method) // just a test to see if the connection is valid
	return n, err
}

func (n *neo4j) send(url string, data string, method string) (string, error) {
	var (
		resp *http.Response // http response
		buf  bytes.Buffer   // contains http response body
		err  error
	)
	if len(url) < 1 {
		url = n.URL + "node" // default path
	}

	client := new(http.Client)
	switch strings.ToLower(method) { // which http method
	case "delete":
		req, e := http.NewRequest("DELETE", url, nil)
		if e != nil {
			err = e
			break
		}
		n.setAuth(*req)
		resp, err = client.Do(req)
	case "post":
		body := strings.NewReader(data)
		req, e := http.NewRequest("POST", url, body)
		if e != nil {
			err = e
			break
		}
		req.Header.Set("Content-Type", "application/json")
		n.setAuth(*req)
		resp, err = client.Do(req)
	case "put":
		body := strings.NewReader(data)
		req, e := http.NewRequest("PUT", url, body)
		if e != nil {
			err = e
			break
		}
		req.Header.Set("Content-Type", "application/json")
		n.setAuth(*req)
		resp, err = client.Do(req)
	case "get":
		fallthrough
	default:
		req, e := http.NewRequest("GET", url, nil)
		if e != nil {
			err = e
			break
		}
		n.setAuth(*req)
		resp, err = client.Do(req)

	}
	if err != nil {
		return "", err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	n.StatusCode = resp.StatusCode // the calling method should do more inspection with chkStatusCode() method and determine if the operation was successful or not.
	return buf.String(), nil
}

func (n *neo4j) setAuth(req http.Request) {
	if len(n.Username) > 0 || len(n.Password) > 0 {
		req.SetBasicAuth(n.Username, n.Password)
	}
}
