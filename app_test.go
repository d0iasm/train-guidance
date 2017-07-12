package app

import (
	"reflect"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSimplerBFS(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()
	from := "a"
	to := "z"
	tracks := []Track{
		{Name: "l1", Stations: []string{"a", "b", "x", "c"}},
		{Name: "l2", Stations: []string{"x", "y", "z"}},
	}
	got := breadthFirstSearch(ctx, from, to, tracks)
	want := []string{"a", "b", "x", "y", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS(%v, %v, %v) = %v but want %v", from, to, tracks, got, want)
	}
}

func TestHarderBFS(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()
	from := "a"
	to := "z"
	tracks := []Track{
		{Name: "l1", Stations: []string{"a", "b", "m", "x", "c"}},
		{Name: "l2", Stations: []string{"x", "y", "z"}},
		{Name: "l3", Stations: []string{"m", "n", "o"}},
	}
	got := breadthFirstSearch(ctx, from, to, tracks)
	want := []string{"a", "b", "x", "y", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS(%v, %v, %v) = %v but want %v", from, to, tracks, got, want)
	}
}
