package web

import (
	"baldeweg/mission/filetypes"
	"baldeweg/mission/logfile"
	"baldeweg/mission/parseJson"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Page struct {
    Title string
    Body []byte
}

var validPath = regexp.MustCompile("^/(list|create|update)$")

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

func Web() {
    http.HandleFunc("/list", makeHandler(listHandler, "GET"))
    http.HandleFunc("/create", makeHandler(createHandler, "POST"))
    http.HandleFunc("/update", makeHandler(updateHandler, "PUT"))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    c := string(parseJson.Write(parseJson.Read(string(logfile.ReadLogfile()))))
    io.WriteString(w, c)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    create := filetypes.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := parseJson.Read(string(logfile.ReadLogfile()))
    t.Missions = append(t.Missions, create)

    logfile.WriteLogfile(parseJson.Write(t))

    c := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    logfile.WriteLogfile(body)

    c := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), method string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method && r.Method != "OPTIONS" {
            http.NotFound(w, r)
            return
        }

        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }

        w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        fn(w, r)
    }
}
