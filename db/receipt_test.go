package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/danyeric123/rewards-api/domain"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open the database: %v", err)
	}
	db.AutoMigrate(&Receipt{})
	db.AutoMigrate(&Item{})
	return db
}

func TestSaveReceipt(t *testing.T) {
	db := setupTestDB(t)
	receiptDB := NewReceiptDB(db)

	testCases := []struct {
		name     string
		receipt  domain.Receipt
		points   int
	}{
		{
			name: "receipt with one item",
			receipt: domain.Receipt{
				Retailer:     "Walmart",
				PurchaseDate: "2021-01-01",
				PurchaseTime: "12:00:00",
				Items: []domain.Item{
					{
						ShortDescription: "Item 1",
						Price:            "1.00",
					},
				},
				Total: "1.00",
			},
			points: 88,
		},
		{
			name: "receipt with multiple items",
			receipt: domain.Receipt{
				Retailer:     "Walmart",
				PurchaseDate: "2021-01-01",
				PurchaseTime: "12:00:00",
				Items: []domain.Item{
					{
						ShortDescription: "Item 1",
						Price:            "1.00",
					},
					{
						ShortDescription: "Item 2",
						Price:            "2.00",
					},
				},
				Total: "3.00",
			},
			points: 176,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := receiptDB.SaveReceipt(tc.receipt, tc.points)
			assert.NoError(t, err)

			receipt, err := receiptDB.GetReceipt(id)
			assert.NoError(t, err)

			assert.Equal(t, tc.receipt.Retailer, receipt.Retailer)
			assert.Equal(t, tc.receipt.PurchaseDate, receipt.PurchaseDate)
			assert.Equal(t, tc.receipt.PurchaseTime, receipt.PurchaseTime)
			assert.Equal(t, tc.receipt.Total, receipt.Total)

			assert.Len(t, receipt.Items, len(tc.receipt.Items))
			for i, item := range receipt.Items {
				assert.Equal(t, tc.receipt.Items[i].ShortDescription, item.ShortDescription)
				assert.Equal(t, tc.receipt.Items[i].Price, item.Price)
			}
		})
	}
}

func TestGetPoints(t *testing.T) {
	db := setupTestDB(t)
	receiptDB := NewReceiptDB(db)

	receipt := Receipt{
		ID: "f47ac10b-58cc-0372-8567-0e02b2c3d479",
		Retailer: "Target",
		PurchaseDate: "2021-01-01",
		PurchaseTime: "12:00:00",
		Total: 12.00,
		Points: 100,
	}

	// Save dummy receipt in database
	err := db.Create(&receipt).Error

	if err != nil {
		t.Fatalf("Failed to save receipt to the database: %v", err)
	}

	testCases := []struct {
		name string
		receiptUUID string
		points int
		err error
	}{
		{
			name: "receipt that exists",
			receiptUUID: receipt.ID,
			points: 100,
			err: nil,
		},
		{
			name: "receipt that does not exist",
			receiptUUID: "f47ac10b-58cc-0372-8567-0e02b2c3d478",
			points: 0,
			err: gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			points, err := receiptDB.GetPoints(tc.receiptUUID)
			assert.Equal(t, tc.points, points)
			assert.Equal(t, tc.err, err)
		})
	}
}
