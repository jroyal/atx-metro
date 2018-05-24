package api

import (
	"net/http"
	"strings"

	"github.com/derekparker/trie"
	"github.com/jroyal/gtfs"
)

func setupStops(stops []*gtfs.Stop) *trie.Trie {
	t := trie.New()
	for _, stop := range stops {
		t.Add(strings.ToLower(stop.Name), stop)
	}
	return t
}

func (s *MetroService) getStops(w http.ResponseWriter, r *http.Request) {
	qArgs := r.URL.Query()
	keys := s.stops.FuzzySearch(strings.ToLower(qArgs.Get("name")))
	stops := []*gtfs.Stop{}
	for _, key := range keys {
		if node, ok := s.stops.Find(key); ok {
			stop := node.Meta().(*gtfs.Stop)
			stops = append(stops, stop)
		}
	}
	render.JSON(w, http.StatusOK, stops)
}
