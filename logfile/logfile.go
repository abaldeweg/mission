package logfile

import (
    "os"
    "log"
    "gopkg.in/yaml.v3"
)

type MissionsLog struct {
    Missions []Mission
}

type Mission struct {
    Date string
    Time string
    Situation string
    Unit string
    Location string
    Links []string
}

func ParseYAML(d []byte) MissionsLog {
    t := MissionsLog{}

    err := yaml.Unmarshal([]byte(d), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return t
}

func WriteYAML(t MissionsLog) []byte {
    d, err := yaml.Marshal(t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return d
}

func IsFile(config string) bool {
    if _, err := os.Stat(config); err == nil {
        return true
    }

    return false
}

func WriteFile(file string, content []byte) {
    d := []byte(content)
    err := os.WriteFile(file, d, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func ReadFile(config string) []byte {
    data, err := os.ReadFile(config)
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
