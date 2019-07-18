package random

import (
	"fmt"
	"math/rand"
	"time"

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
	// ...

	array := []int{1, 3, 4, 5, 6, 7, 8, 9}
	array2 := [][]int{{2, 3}, {8, 9}, {11, 12}, {14, 15, 16, 29, 30}, {17, 18}, {20, 21}, {23, 24, 25, 31, 32}, {26, 27, 28, 36, 34}}
	array3 := []int{1, 2, 3, 4, 5}

	headerID := make([]int, 0)
	headerID = append(array)

	fieldId := make([][]int, 0)
	fieldId = append(array2)

	alertType := make([]int, 0)
	alertType = append(array3)

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	arr1Index := rand.Intn(len(headerID))
	//res := array2[arr1Index]
	arr2Index := rand.Intn(len(fieldId[arr1Index]))
	fmt.Println(headerID[arr1Index])
	//fmt.Println(res)
	fmt.Println(fieldId[arr1Index][arr2Index])

	arr3Index := rand.Intn(len(alertType))
	fmt.Println(arr3Index)

	context.SetOutput("headerID", headerID[arr1Index])
	context.SetOutput("fieldID", fieldId[arr1Index][arr2Index])
	context.SetOutput("alertType", arr3Index)

	return true, nil
}
