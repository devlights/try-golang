package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	appLog = log.New(os.Stdout, "", 0)
	tmpls  = map[string]*template.Template{
		"index": template.Must(template.ParseFiles("tmpl/index.html")),
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			now = time.Now().Format(time.RFC3339)
		)
		appLog.Printf("%s: %s\n", now, r.URL)

		if t, ok := tmpls["index"]; ok {
			t.Execute(w, struct{ Now string }{now})
		}
	})
	http.HandleFunc("/notfound", http.NotFound)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
