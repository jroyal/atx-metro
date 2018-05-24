package api

import (
	"encoding/json"
	"net/http"
)

func (s *MetroService) getAgency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.feed.Agencies)
}
