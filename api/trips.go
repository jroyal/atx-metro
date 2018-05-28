package api

import (
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

func setupTrips(trips []*gtfs.Trip) map[string]*gtfs.Trip {
	tripMap := make(map[string]*gtfs.Trip)
	for _, trip := range trips {
		tripMap[trip.ID] = trip
	}
	return tripMap
}

func setupStopToRoutes() {

}

func (s *MetroService) getTrips(w http.ResponseWriter, r *http.Request) {
	routeID := chi.URLParam(r, "routeID")
	render.JSON(w, http.StatusOK, s.routeToTripsMap[routeID])
}
