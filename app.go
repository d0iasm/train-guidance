package app

import (
	// "fmt"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Track struct {
	Name     string   `json:"Name"`
	Stations []string `json:"Stations"`
}

func init() {
	http.Handle("./stylesheets/", http.StripPrefix("./stylesheets/", http.FileServer(http.Dir("./stylesheets/"))))
	http.HandleFunc("/", handleSearch)
	http.HandleFunc("/result", handleResult)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	tracks := getTracks(w, r)
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, tracks)
}

func handleResult(w http.ResponseWriter, r *http.Request) {
	tracks := getTracks(w, r)

	from := r.FormValue("from")
	to := r.FormValue("to")
	result := breadthFirstSeatch(from, to, tracks)

	tmpl := template.Must(template.ParseFiles("./result.html"))
	tmpl.Execute(w, result)
}

func getTracks(w http.ResponseWriter, r *http.Request) []Track {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://tokyo.fantasy-transit.appspot.com/net?format=json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var tracks []Track
	if err := json.Unmarshal(body, &tracks); err != nil {
		return nil
	}

	return tracks
}

func breadthFirstSeatch(from, to string, tracks []Track) []string {
	path := []string{}
	for _, track := range tracks {
		path = append(path, string(track.Name))
	}
	return path
}
