package cache

import (
	"fetch-demo/internal/api"
	"fetch-demo/internal/process"
	"sync"

	"github.com/google/uuid"
)

// RecieptCache represents the cache entry for a receipt and its associated points.
type RecieptCache struct {
	Reciept api.Receipt
	Points  int
}

var (
	// receiptMap stores the cached receipts with their corresponding UUID as the key.
	receiptMap = make(map[uuid.UUID]RecieptCache)

	// mu is used to synchronize access to the receiptMap.
	mu sync.Mutex
)

// AddToCache adds a receipt to the cache and calculates its points.
//
// It processes the receipt to determine the points, stores the receipt in the cache
// with a unique UUID, and returns the UUID and any error encountered during processing.
func AddToCache(reciept api.Receipt) (uuid.UUID, error) {
	mu.Lock()
	defer mu.Unlock()

	cacheID := uuid.New()
	points, err := process.ProcessReciept(reciept)
	if err != nil {
		return uuid.Nil, err
	}
	receiptMap[cacheID] = RecieptCache{Reciept: reciept, Points: points}
	return cacheID, nil
}

// GetRecord retrieves a receipt cache entry by its UUID.
//
// It returns the cached receipt entry if found, or nil if the record does not exist.
func GetRecord(cacheID uuid.UUID) *RecieptCache {
	mu.Lock()
	defer mu.Unlock()

	if entry, exists := receiptMap[cacheID]; exists {
		return &entry
	}
	return nil
}
