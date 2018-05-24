package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jroyal/gtfs"
)

func setupRouteToTrips(trips []*gtfs.Trip) map[string][]*gtfs.Trip {
	tripMap := make(map[string][]*gtfs.Trip)
	for _, trip := range trips {
		tripMap[trip.RouteID] = append(tripMap[trip.RouteID], trip)
	}
	return tripMap
}

func (s *MetroService) getTrips(w http.ResponseWriter, r *http.Request) {
	routeID := chi.URLParam(r, "routeID")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.routeToTripsMap[routeID])
}
