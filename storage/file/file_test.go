package file

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
    msg := GetUrl()

    if reflect.TypeOf(msg).String() != "string" {
        t.Fatalf("%s is not type of %s", reflect.TypeOf(msg),  "string")
    }
}
