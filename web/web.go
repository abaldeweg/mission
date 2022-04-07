package web

import (
	"baldeweg/mission/commands/create"
	"baldeweg/mission/logfile"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
    Title string
    Body []byte
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/list.html"))
var validPath = regexp.MustCompile("^/(list|create|edit|update)$")

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

func Web() {
    http.HandleFunc("/list", makeHandler(listHandler))
    http.HandleFunc("/create", makeHandler(createHandler))
    http.HandleFunc("/edit", makeHandler(editHandler))
    http.HandleFunc("/update", makeHandler(updateHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{"List", logfile.ReadLogfile()}
    renderTemplate(w, "list", p)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    create.Create()
    http.Redirect(w, r, "/list", http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{"Edit", logfile.ReadLogfile()}
    renderTemplate(w, "edit", p)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body := r.FormValue("body")
    p := &Page{Title: "Update", Body: []byte(body)}
    logfile.WriteLogfile(p.Body)
    http.Redirect(w, r, "/list", http.StatusFound)
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
