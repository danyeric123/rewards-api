package domain

import (
	"math"
	"unicode"

	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
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
	if math.Mod(receiptTotal, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items purchased
	points += len(r.Items) / 2 * 5 // integer division rounds down

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range r.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := item.GetPrice()
			if err != nil {
				logrus.WithError(err).Error("Failed to get price from item while calculating points")
				return 0, err
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is an odd number
	// Assuming the date is in the format YYYY-MM-DD
	date := strings.Split(r.PurchaseDate, "-")
	day, err := strconv.Atoi(date[2])
	if err != nil {
		logrus.WithError(err).Error("Failed to convert day to integer")
		return 0, err
	}
	if day%2 != 0 {
		points += 6
	}

	// 10 points if the purchase time is between 2:00 PM and 4:00 PM
	// Assuming the time is in the format HH:MM and the time is in 24-hour format
	// and the range is inclusive not exclusive
	time := strings.Split(r.PurchaseTime, ":")
	hour, err := strconv.Atoi(time[0])
	if err != nil {
		logrus.WithError(err).Error("Failed to convert hour to integer")
		return 0, err
	}
	if hour >= 14 && hour <= 16 {
		points += 10
	}

	logrus.Info("Points calculated: ", points)
	return points, nil
}
