package api

import (
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
	render.JSON(w, http.StatusOK, s.feed.Routes)
}

func (s *MetroService) getRoute(w http.ResponseWriter, r *http.Request) {
	routeID := chi.URLParam(r, "id")
	render.JSON(w, http.StatusOK, s.routeMap[routeID])
}
