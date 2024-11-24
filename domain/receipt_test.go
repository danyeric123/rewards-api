package domain

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected int
	}{
		{
			name: "Receipt with two items and whole dollar total",
			receipt: Receipt{
				Retailer:     "Retailer1",
				PurchaseDate: "2023-10-01",
				PurchaseTime: "14:30",
				Items: []Item{
					{ShortDescription: "Item1", Price: "1.00"},
					{ShortDescription: "Item2", Price: "2.00"},
				},
				Total: "3.00",
			},
			// 9 (Retailer1 = len 9) + 50 (round dollar) 
			// + 25 (multiple of 0.25) + 5 (2 items) + 6 (odd day) + 10 (time)
			expected: 105,
		},
		{
			name: "No points for even day",
			receipt: Receipt{
				Retailer:     "Retailer2",
				PurchaseDate: "2023-10-02",
				PurchaseTime: "14:30",
				Items: []Item{
					{ShortDescription: "Item1", Price: "1.00"},
					{ShortDescription: "Item2", Price: "2.00"},
				},
				Total: "3.00",
			},
			// 9 (Retailer2 = len 9) + 50 (round dollar) + 25 (multiple of 0.25) + 5 (2 items) + 10 (time)
			expected: 99,
		},
		{
			name: "Points for item description length multiple of 3",
			receipt: Receipt{
				Retailer:     "Retailer3",
				PurchaseDate: "2023-10-03",
				PurchaseTime: "15:00",
				Items: []Item{
					{ShortDescription: "Item12", Price: "15.00"},
					{ShortDescription: "Item2", Price: "2.00"},
				},
				Total: "3.00",
			},
			// 9 (Retailer3 = len 9) + 50 (round dollar) + 25 (multiple of 0.25) + 5 (2 items) 
			// + 6 (odd day) + 10 (time) + 3 (item description length multiple of 3 so multiply price by 0.2 and round up)
			expected: 108,
		},
		{
			name: "Points for item description length multiple of 3 with cents",
			receipt: Receipt{
				Retailer:     "Retailer4",
				PurchaseDate: "2023-10-04",
				PurchaseTime: "15:00",
				Items: []Item{
					{ShortDescription: "Item12", Price: "10.00"},
					{ShortDescription: "Item2", Price: "2.00"},
				},
				Total: "3.50",
			},
			// 9 (Retailer4 = len 9) + 25 (multiple of 0.25) + 5 (2 items) + 10 (time)
			// + 2 (item description for first item is length multiple of 3 so multiply price by 0.2 and round up)
			expected: 51,
		},
		{
			name: "Points for item description length not multiple of 3",
			receipt: Receipt{
				Retailer:     "Retailer5",
				PurchaseDate: "2023-10-05",
				PurchaseTime: "15:00",
				Items: []Item{
					{ShortDescription: "Item12", Price: "1.00"},
					{ShortDescription: "Item2", Price: "2.00"},
				},
				Total: "3.00",
			},
			// 9 (Retailer5 = len 9) + 50 (round dollar) 
			// + 25 (multiple of 0.25) + 5 (2 items) + 6 (odd day) + 10 (time)
			expected: 105,
		}, 
		{
			name: "Points for two items with length multiple of 3",
			receipt: Receipt{
				Retailer:     "Retailer6",
				PurchaseDate: "2023-10-06",
				PurchaseTime: "15:00",
				Items: []Item{
					{ShortDescription: "Item12", Price: "15.00"},
					{ShortDescription: "Item12", Price: "20.00"},
				},
				Total: "3.00",
			},
			// 9 (Retailer6 = len 9) + 50 (round dollar)
			// + 25 (multiple of 0.25) + 5 (2 items) + 6 (odd day) + 10 (time)
			// + 5 (two item description for both items is length multiple of 3 so multiply price by 0.2 and round up)
			expected: 106,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			points, err := tt.receipt.CalculatePoints()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, points)
		})
	}
}

func TestGetTotal(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected float64
		expectedError		error
	}{
		{
			name: "Valid total",
			receipt: Receipt{
				Total: "123.45",
			},
			expected: 123.45,
			expectedError: nil,
		},
		{
			name: "Invalid total",
			receipt: Receipt{
				Total: "invalid",
			},
			expected: 0,
			expectedError: &strconv.NumError{
				Func: "ParseFloat",
				Num:  "invalid",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			total, err := tt.receipt.GetTotal()
			if tt.expected == 0 {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, total)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, total)
			}
		})
	}
}

func TestGetPrice(t *testing.T) {
	tests := []struct {
		name     string
		item     Item
		expected float64
		expectedError 	 error
	}{
		{
			name: "Valid price",
			item: Item{
				Price: "123.45",
			},
			expected: 123.45,
			expectedError: nil,
		},
		{
			name: "Invalid price",
			item: Item{
				Price: "invalid",
			},
			expected: 0,
			expectedError: &strconv.NumError{
				Func: "ParseFloat",
				Num:  "invalid",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			price, err := tt.item.GetPrice()
			if tt.expected == 0 {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, price)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, price)
			}
		})
	}
}