package logfile

import (
    "testing"
    "reflect"
)

func TestConfig(t *testing.T) {
    msg := Path()

    if reflect.TypeOf(msg).String() != "string" {
        t.Fatalf("%s is not type of %s", reflect.TypeOf(msg),  "string")
    }
}
