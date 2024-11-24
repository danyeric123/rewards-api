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
