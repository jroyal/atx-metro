package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jroyal/gtfs"
)

func setupRoutes(routes []*gtfs.Route) map[string]*gtfs.Route {
	routeMap := make(map[string]*gtfs.Route)
	for _, route := range routes {
		routeMap[route.ID] = route
	}
	return routeMap
}

func (s *MetroService) getRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.feed.Routes)
}

func (s *MetroService) getRoute(w http.ResponseWriter, r *http.Request) {
	routeID := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.routeMap[routeID])
}
