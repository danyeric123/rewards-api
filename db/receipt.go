package db

import (
	"github.com/danyeric123/rewards-api/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReceiptDB struct {
	DB *gorm.DB
}

func NewReceiptDB(db *gorm.DB) *ReceiptDB {
	return &ReceiptDB{DB: db}
}

// SaveReceipt saves a receipt to the database
// and returns the ID of the saved receipt
func (r *ReceiptDB) SaveReceipt(receipt domain.Receipt, points int) (string, error) {
	// Having a UUID in a SQL database is not good for performance
	// because it's a random value; it would be better to use a sequential ID
	// for the primary key. There could also be collisions if the UUID is not
	// unique and then you would have to retry the operation multiple times,
	// thus making multiple queries when you could have done it in one.
	// This, though, was part of the prompt; hence for simplicity, I am assuming
	// that the UUID is unique.
	receiptID := uuid.New().String()

	receiptModel := Receipt{
		ID:           receiptID,
		Retailer:     receipt.Retailer,
		PurchaseDate: receipt.PurchaseDate,
		PurchaseTime: receipt.PurchaseTime,
		Total:        receipt.Total,
		Points:       points,
	}

	err := r.DB.Create(&receiptModel).Error
	if err != nil {
		logrus.WithError(err).Error("Failed to save receipt to the database")
		if err := r.DB.Rollback().Error; err != nil {
			logrus.WithError(err).Error("Failed to rollback transaction")
		}
		return "", err
	}
	
	return receiptID, nil
}

// GetPoints returns the points for a receipt
func (r *ReceiptDB) GetPoints(id string) (int, error) {
	// Get the points for the receipt from the database
	return 0, nil
}
