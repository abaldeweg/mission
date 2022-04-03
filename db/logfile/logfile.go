package logfile

import (
	"log"
	"os"
	"path"

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

var dir string

func init() {
    log.SetPrefix("logfile: ")
    log.SetFlags(0)
}

func SetPath(d string) {
    _, err := os.Stat(d)
    if err != nil {
        log.Fatal(err)
    }
    if err == nil {
        dir = d
    }
}

func GetPath() string {
    return dir
}

func GetUrl() string {
    return path.Join(GetPath(), "missions.yaml")
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

func HasLogfile() bool {
    if _, err := os.Stat(GetUrl()); err == nil {
        return true
    }

    return false
}

func WriteLogfile(content []byte) {
    err := os.WriteFile(GetUrl(), content, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func ReadLogfile() []byte {
    data, err := os.ReadFile(GetUrl())
    if err != nil {
        log.Fatal(err)
    }

    return data
}

func CreateTemplate() {
    WriteLogfile(WriteYAML(ParseYAML([]byte(data))))
}
