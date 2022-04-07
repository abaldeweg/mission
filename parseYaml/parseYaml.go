package parseYaml

import (
	"baldeweg/mission/filetypes"
	"log"

	"gopkg.in/yaml.v3"
)

var data = `
notes:
replacements:
missions:
`

func init() {
    log.SetPrefix("parseYaml: ")
    log.SetFlags(0)
}

func ParseYAML(d []byte) filetypes.Logfile {
    t := filetypes.Logfile{}

    err := yaml.Unmarshal([]byte(d), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return t
}

func WriteYAML(t filetypes.Logfile) []byte {
    d, err := yaml.Marshal(t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return d
}

func Template() []byte {
    return WriteYAML(ParseYAML([]byte(data)))
}
