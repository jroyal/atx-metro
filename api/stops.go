package api

import (
	"net/http"
	"strings"

	"github.com/derekparker/trie"
	"github.com/jroyal/gtfs"
)

type ExpandedStop struct {
	*gtfs.Stop
	Routes []string
}

func setupStopTimes(stopTimes []*gtfs.StopTimes) map[string][]*gtfs.StopTimes {
	stoptimeMap := make(map[string][]*gtfs.StopTimes)
	for _, stopTime := range stopTimes {
		stoptimeMap[stopTime.StopID] = append(stoptimeMap[stopTime.StopID], stopTime)
	}
	return stoptimeMap
}

func setupStops(stops []*gtfs.Stop, trips map[string]*gtfs.Trip, stopTimes map[string][]*gtfs.StopTimes) *trie.Trie {
	t := trie.New()
	for _, stop := range stops {
		times := stopTimes[stop.ID]
		routeMap := map[string]bool{}
		for _, time := range times {
			routeMap[trips[time.TripID].RouteID] = true
		}
		routes := []string{}
		for k := range routeMap {
			routes = append(routes, k)
		}
		t.Add(strings.ToLower(stop.Name), ExpandedStop{stop, routes})
	}
	return t
}

func (s *MetroService) getStops(w http.ResponseWriter, r *http.Request) {
	qArgs := r.URL.Query()
	nameFilter := strings.ToLower(qArgs.Get("name"))
	if nameFilter == "" {
		render.JSON(w, http.StatusBadRequest, ApiError{"name is required"})
		return
	}
	keys := s.stops.FuzzySearch(nameFilter)
	stops := []ExpandedStop{}
	for _, key := range keys {
		if node, ok := s.stops.Find(key); ok {
			stop := node.Meta().(ExpandedStop)
			stops = append(stops, stop)
		}
	}
	render.JSON(w, http.StatusOK, stops)
}
