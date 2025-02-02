package process

import (
	"fetch-demo/internal/api"
	"testing"
	"time"

	"github.com/oapi-codegen/runtime/types"
)

// Test ProcessReciept function
func TestProcessReciept(t *testing.T) {
	tests := []struct {
		name          string
		receipt       api.Receipt
		expectedPts   int
		expectedError bool
	}{
		{
			name: "Valid receipt",
			receipt: api.Receipt{
				Retailer:     "Target",
				PurchaseDate: parseDate("2022-01-01"),
				PurchaseTime: "13:01",
				Items: []api.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expectedPts:   28,
			expectedError: false,
		},
		{
			name: "Invalid total value",
			receipt: api.Receipt{
				Retailer:     "Target",
				PurchaseDate: parseDate("2022-01-01"),
				PurchaseTime: "13:01",
				Items: []api.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "invalid",
			},
			expectedPts:   0,
			expectedError: true,
		},
		{
			name: "Invalid item price",
			receipt: api.Receipt{
				Retailer:     "Target",
				PurchaseDate: parseDate("2022-01-01"),
				PurchaseTime: "13:01",
				Items: []api.Item{
					{ShortDescription: "Invalid Price Item", Price: "invalid"},
				},
				Total: "35.35",
			},
			expectedPts:   0,
			expectedError: true,
		},
		{
			name: "Empty retailer name",
			receipt: api.Receipt{
				Retailer:     "",
				PurchaseDate: parseDate("2022-01-01"),
				PurchaseTime: "13:01",
				Items: []api.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "35.35",
			},
			expectedPts:   6,
			expectedError: false,
		},
		{
			name: "Whole number total (50 points)",
			receipt: api.Receipt{
				Retailer:     "Walmart",
				PurchaseDate: parseDate("2022-01-02"),
				PurchaseTime: "10:00",
				Items: []api.Item{
					{ShortDescription: "Item 1", Price: "10.00"},
					{ShortDescription: "Item 2", Price: "5.00"},
				},
				Total: "15.00",
			},
			expectedPts:   90, // 7 + 5 + 2 + 1 + 25 + 50
			expectedError: false,
		},
		{
			name: "Total is a multiple of 0.25 (25 points)",
			receipt: api.Receipt{
				Retailer:     "BestBuy",
				PurchaseDate: parseDate("2022-01-03"),
				PurchaseTime: "15:00",
				Items: []api.Item{
					{ShortDescription: "Item 1", Price: "8.00"},
					{ShortDescription: "Item 2", Price: "7.25"},
				},
				Total: "15.25",
			},
			expectedPts:   57, // 7 + 6 + 10 + 5 + 25 + 2 + 2
			expectedError: false,
		},
		{
			name: "Receipt with multiple Gatorade items",
			receipt: api.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: parseDate("2022-03-20"),
				PurchaseTime: "14:33",
				Items: []api.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expectedPts:   109,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPts, err := ProcessReciept(tt.receipt)
			if (err != nil) != tt.expectedError {
				t.Errorf("ProcessReciept() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if gotPts != tt.expectedPts {
				t.Errorf("ProcessReciept() = %v, want %v", gotPts, tt.expectedPts)
			}
		})
	}
}

// Helper function to parse date string into openapi_types.Date
func parseDate(dateStr string) types.Date {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Return a Date instance with the parsed time
	return types.Date{Time: parsedDate}
}
