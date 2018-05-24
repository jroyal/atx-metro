package api

import (
	"net/http"

	"github.com/derekparker/trie"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jroyal/gtfs"
)

type MetroService struct {
	Handler http.Handler
	feed    gtfs.Feed

	routeMap        map[string]*gtfs.Route
	routeToTripsMap map[string][]*gtfs.Trip
	stops           *trie.Trie
}

type MetroServiceConfig struct {
	FeedDirectory string
}

func NewService(config *MetroServiceConfig) *MetroService {
	feed := gtfs.LoadFromDirectory(config.FeedDirectory)
	service := &MetroService{
		feed:            feed,
		routeMap:        setupRoutes(feed.Routes),
		routeToTripsMap: setupRouteToTrips(feed.Trips),
		stops:           setupStops(feed.Stops),
	}
	api := chi.NewRouter()
	api.Get("/agency", service.getAgency)
	api.Get("/routes", service.getRoutes)
	api.Get("/routes/{id}", service.getRoute)
	api.Get("/trips/{routeID}", service.getTrips)
	api.Get("/stops", service.getStops)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Mount("/api", api)
	service.Handler = router
	return service
}
