package parseJson

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
    blob := `{"notes":[],"replacements":{},"missions":[]}`
	unmarshal := Read(blob)
    if reflect.TypeOf(unmarshal).String() != "parseJson.Logfile" {
        t.Fatalf("%s is not type of %s", reflect.TypeOf(unmarshal),  "parsejson.Logfile")
    }

    marshal := Write(unmarshal)
    if fmt.Sprintf("%s", reflect.ValueOf(marshal)) != `{"Notes":[],"Replacements":{},"Missions":[]}` {
        t.Fatalf("%s is not type of %s", reflect.ValueOf(marshal),  "[]uint8")
    }
}
