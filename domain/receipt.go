package domain

import (
	"math"
	"unicode"

	"github.com/sirupsen/logrus"
	"strconv"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	// TODO: This comes in as a string, but we want to store it as a float64
	Total string `json:"total"`
}

func (r *Receipt) GetTotal() (float64, error) {
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to convert total to float64")
		return 0, err
	}
	return total, nil
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	// TODO: This comes in as a string, but we want to store it as a float64
	Price string `json:"price"`
}

func (i *Item) GetPrice() (float64, error) {
	price, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to convert price to float64")
		return 0, err
	}
	return price, nil
}

type ProcessResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type GetResponse struct {
	Points int `json:"points"`
}

func (r *Receipt) CalculatePoints() (int, error) {
	logrus.Info("Calculating points for receipt")
	points := 0
	// Calculate points based on rules

	// 1 point for every alphanumeric character in the retailer name
	for _, c := range r.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			points++
		}
	}

	receiptTotal, err := r.GetTotal()
	if err != nil {
		return 0, err
	}

	// 50 points if total is a round dollar amount with no cents
	if math.Mod(receiptTotal, 1.0) == 0 {
		points += 50
	}

	// 25 points if total is a multiple of 0.25

	// 5 points for every two items purchased

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.

	// 6 points if the day in the purchase date is an odd number

	// 10 points if the purchase time is between 2:00 PM and 4:00 PM
	logrus.Info("Points calculated: ", points)
	return points, nil
}
