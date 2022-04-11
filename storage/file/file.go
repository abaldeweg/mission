package file

import (
	"log"
	"os"
	"path"
)

var dir string

func init() {
    log.SetPrefix("file: ")
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
    return path.Join(GetPath(), "missions.json")
}

func Exists(filename string) bool {
    if _, err := os.Stat(GetUrl()); err == nil {
        return true
    }

    return false
}

func Write(filename string, content []byte) {
    err := os.WriteFile(GetUrl(), content, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func Read(filename string) []byte {
    data, err := os.ReadFile(GetUrl())
    if err != nil {
        log.Fatal(err)
    }
    return data
}
