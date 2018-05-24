package api

import (
	"net/http"
)

func (s *MetroService) getAgency(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, http.StatusOK, s.feed.Agencies)
}
