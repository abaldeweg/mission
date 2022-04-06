package util

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Logfile struct {
    Notes []string
    Replacements map[string]string
    Missions []Mission
}

type Mission struct {
    Date string
    Time string
    Keyword string
    Situation string
    Unit string
    Location string
    Links []string
}

var data = `
notes:
replacements:
missions:
`

func init() {
    log.SetPrefix("util: ")
    log.SetFlags(0)
}

func ParseYAML(d []byte) Logfile {
    t := Logfile{}

    err := yaml.Unmarshal([]byte(d), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return t
}

func WriteYAML(t Logfile) []byte {
    d, err := yaml.Marshal(t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return d
}

func Template() []byte {
    return WriteYAML(ParseYAML([]byte(data)))
}
