package restapi

import (
	"encoding/json"
	"fetch-demo/internal/api"
	"fetch-demo/internal/cache"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Server represents the HTTP server handling receipt processing.
type Server struct{}

// NewServer creates and returns a new instance of the Server.
func NewServer() Server {
	return Server{}
}

// PostReceiptsProcess handles receipt submission.
//
// It decodes the request body into a Receipt object, adds the receipt to cache,
// and returns a unique ID for the stored receipt. It also logs the receipt details.
func (s *Server) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	var receipt api.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := cache.AddToCache(receipt)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"id": id.String(),
	}

	log.Printf("Received receipt from retailer: %s, total: %s", receipt.Retailer, receipt.Total)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetReceiptsIdPoints returns points for a given receipt ID.
//
// It retrieves the receipt record from the cache using the provided UUID and
// returns the points associated with the receipt. If the record is not found,
// it responds with a 404 error.
func (s *Server) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	parseUUID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid UUID input", http.StatusNotFound)
		return
	}

	recieptRec := cache.GetRecord(parseUUID)
	if recieptRec == nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
	response := map[string]interface{}{
		"points": recieptRec.Points,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
