package logfile

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Logfile struct {
    Vars map[string]string
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

func init() {
    log.SetPrefix("logfile: ")
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

func HasLogfile() bool {
    if _, err := os.Stat(Path()); err == nil {
        return true
    }

    return false
}

func WriteLogfile(content []byte) {
    err := os.WriteFile(Path(), content, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func ReadLogfile() []byte {
    data, err := os.ReadFile(Path())
    if err != nil {
        log.Fatal(err)
    }

    return data
}

func Path() string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    return home + "/missions.yaml"
}
