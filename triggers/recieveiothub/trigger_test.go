package recieveiothub

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "id": "mytrigger",
  "settings": {
    "setting": "somevalue"
  },
  "handlers": [
    {
      "settings": {
        "handler_setting": "somevalue"
      },
      "action" {
	     "id": "test_action"
      }
    }
  ]
}`

func TestCreate(t *testing.T) {

	// New factory
<<<<<<< HEAD
	md := trigger.NewMetadata(getJSONMetadata())
=======
	md := trigger.NewMetadata(getJsonMetadata())
>>>>>>> f7c5b9dd207fc52963987031e351190924041fce
	f := NewFactory(md)

	if f == nil {
		t.Fail()
	}

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	trg := f.New(&config)

	if trg == nil {
		t.Fail()
	}
}
