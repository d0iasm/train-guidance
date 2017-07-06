package app

import (
	// "fmt"
	"container/list"
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

func makeAdjacentMap(tracks []Track) map[string][]string {
	adjacent := map[string][]string{}
	for _, track := range tracks {
		for i, station := range track.Stations {
			if i > 0 {
				adjacent[station] = append(adjacent[station], string(track.Stations[i-1]))
			}
			if i < len(track.Stations)-1 {
				adjacent[station] = append(adjacent[station], string(track.Stations[i+1]))
			}
		}
	}
	return adjacent
}

func breadthFirstSearch(from, to string, tracks []Track) []string {
	// path := []string{}
	adjacent := makeAdjacentMap(tracks)
	queue := list.New()
	queue.PushBack([]string{from})
	current := ""
	path := []string{}
	for queue.Len() > 0 {
		path, _ = queue.Remove(queue.Front()).([]string)
		// path = append(path, current)
		current = path[len(path)-1]
		if current == to {
			return path
		}
		for _, nextStation := range adjacent[current] {
			queue.PushBack(append(path, nextStation))
		}
	}
	return nil
}
