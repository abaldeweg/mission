package logfile

import (
	"baldeweg/mission/db/bucket"
	"fmt"
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
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        if _, err := os.Stat(GetUrl()); err == nil {
            return true
        }

        return false
    }

    if storage == "bucket" {
        return bucket.Exists(os.Getenv("BUCKET_NAME"), "missions.yaml")
    }

    return false
}

func WriteLogfile(content []byte) {
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        err := os.WriteFile(GetUrl(), content, 0644)
        if err != nil {
            log.Fatal(err)
        }
    }

    if storage == "bucket" {
        bucket.Write(os.Getenv("BUCKET_NAME"), "missions.yaml", fmt.Sprintf("%s", content))
    }
}

func ReadLogfile() []byte {
    var d []byte
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        data, err := os.ReadFile(GetUrl())
        if err != nil {
            log.Fatal(err)
        }
        d = data
    }

    if storage == "bucket" {
        d = bucket.Read(os.Getenv("BUCKET_NAME"), "missions.yaml")
    }

    return d
}

func CreateTemplate() {
    WriteLogfile(WriteYAML(ParseYAML([]byte(data))))
}
