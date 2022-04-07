package filetypes

import "log"

func init() {
    log.SetPrefix("filetypes: ")
    log.SetFlags(0)
}

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

type Msg struct {
    Msg string
}
