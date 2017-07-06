package app

import (
	"html/template"
	"net/http"
	// "google.golang.org/appengine"
	// "google.golang.org/appengine/urlfetch"
)

func init() {
	http.Handle("./stylesheets/", http.StripPrefix("./stylesheets/", http.FileServer(http.Dir("./stylesheets/"))))
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	// client := urlfetch.Client(ctx)
	// resp, err := client.Get("http://tokyo.fantasy-transit.appspot.com/net?format=json")
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// return
	// }
	// fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)

	a := r.FormValue("a")
	b := r.FormValue("b")
	result := combine(a, b)
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, result)
}

func combine(a, b string) string {
	return a + b
}
