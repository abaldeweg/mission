package web

import (
	"baldeweg/mission/commands/create"
	"baldeweg/mission/filetypes"
	"baldeweg/mission/logfile"
	"baldeweg/mission/parseJson"
	"baldeweg/mission/parseYaml"
	"io"
	"log"
	"net/http"
	"regexp"
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
    http.HandleFunc("/list", makeHandler(listHandler))
    http.HandleFunc("/create", makeHandler(createHandler))
    http.HandleFunc("/update", makeHandler(updateHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    p := string(parseJson.Write(parseYaml.ParseYAML(logfile.ReadLogfile())))
    io.WriteString(w, p)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    create.Create()

    p := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, p)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    p := &Page{Title: "Update", Body: body}
    logfile.WriteLogfile(p.Body)

    c := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r)
    }
}
