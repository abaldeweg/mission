package logfile

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Logfile struct {
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

var path string

func init() {
    log.SetPrefix("logfile: ")
    log.SetFlags(0)
}

func SetPath(dir string) {
    _, err := os.Stat(dir)
    if err != nil {
        log.Fatal(err)
    }
    if err == nil {
        path = dir
    }
}

func GetPath() string {
    return path
}

func GetUrl() string {
    return GetPath() + "/missions.yaml"
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
