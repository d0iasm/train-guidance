package app

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"html/template"
	"net/http"
)

func init() {
	http.Handle("./stylesheets/", http.StripPrefix("./stylesheets/", http.FileServer(http.Dir("./stylesheets/"))))
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	// client := urlfetch.Client(ctx)
	// resp, err := client.Get("http://tokyo.fantasy-transit.appspot.com/net?format=json")
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// return
	// }
	// fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://www.google.com/")

	from := r.FormValue("from")
	to := r.FormValue("to")
	result := combine(from, to)
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, result)
}

func combine(a, b string) string {
	return a + b
}
