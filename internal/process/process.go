package process

import (
	"fetch-demo/internal/api"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// ProcessReciept calculates the total points earned from a receipt.
//
// It adds points based on various factors such as retailer name length,
// receipt total, items purchased, and purchase date/time.
func ProcessReciept(reciept api.Receipt) (int, error) {
	totalPts := 0

	// Name points
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	totalPts += len(re.ReplaceAllString(strings.TrimSpace(reciept.Retailer), ""))

	// Reciept Total Points
	recTotal, err := processTotal(reciept)
	if err != nil {
		return 0, err
	}
	totalPts += recTotal

	// Item Points
	itemPts, err := processItems(reciept)
	if err != nil {
		return 0, err
	}
	totalPts += itemPts

	// Date/Time Points
	dtPoints, err := processDateTime(reciept)
	if err != nil {
		return 0, err
	}
	totalPts += dtPoints

	return totalPts, nil
}

// processTotal calculates points based on the receipt's total amount.
//
// Points are awarded based on whether the total is a whole number or a
// multiple of 0.25.
func processTotal(reciept api.Receipt) (int, error) {
	totalPts := 0
	// Total Prep
	decTotal, err := decimal.NewFromString(reciept.Total)
	if err != nil {
		log.Print("Total is not a number.")
		return 0, err
	}

	// Whole num points
	if decTotal.IsInteger() {
		totalPts += 50
	}

	// Multiple of 0.25
	qtr, _ := decimal.NewFromString("0.25")
	if decTotal.Mod(qtr).Equal(decimal.Zero) {
		totalPts += 25
	}
	return totalPts, nil
}

// processItems calculates points based on the items in the receipt.
//
// Points are awarded based on the number of items and specific conditions
// on the item descriptions and prices.
func processItems(reciept api.Receipt) (int, error) {
	totalPts := 0

	// 5 pts for every two items
	totalPts += 5 * (len(reciept.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	ptMult, _ := decimal.NewFromString("0.2")
	for _, item := range reciept.Items {

		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := decimal.NewFromString(item.Price)
			if err != nil {
				log.Print("Item price is not a number.")
				return 0, err
			}
			totalPts += int(price.Mul(ptMult).RoundUp(0).IntPart())
		}
	}
	return totalPts, nil
}

// processDateTime calculates points based on the receipt's purchase date and time.
//
// Points are awarded if the purchase date's day is odd or if the time of purchase
// falls between 2:00pm and 4:00pm.
func processDateTime(reciept api.Receipt) (int, error) {
	totalPts := 0

	// 6 points if the day in the purchase date is odd.
	if reciept.PurchaseDate.Day()%2 != 0 {
		totalPts += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timestamp, err := time.Parse("15:04", reciept.PurchaseTime)
	if err != nil {
		log.Print("Timestmp is not a valid time.")
		return 0, err
	}
	tsHour := timestamp.Hour()
	if tsHour >= 14 && tsHour < 16 {
		totalPts += 10
	}
	return totalPts, nil
}
