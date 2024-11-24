package db

import (
	"github.com/danyeric123/rewards-api/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// TODO: Consider changing to transactions with rollbacks in case of errors

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

	receiptTotal, err := receipt.GetTotal()

	if err != nil {
		logrus.WithError(err).Error("Failed to get total from receipt")
		return "", err
	}

	receiptModel := Receipt{
		ID:           receiptID,
		Retailer:     receipt.Retailer,
		PurchaseDate: receipt.PurchaseDate,
		PurchaseTime: receipt.PurchaseTime,
		Total:        receiptTotal,
		Points:       points,
	}

	err = r.DB.Create(&receiptModel).Error
	if err != nil {
		logrus.WithError(err).Error("Failed to save receipt to the database")
		return "", err
	}
	var items []Item
	for _, item := range receipt.Items {
		itemPrice, err := item.GetPrice()
		if err != nil {
			logrus.WithError(err).Error("Failed to get price from item")
			return "", err
		}
		itemModel := Item{
			ID:               uuid.New().String(),
			ReceiptID:        receiptID,
			ShortDescription: item.ShortDescription,
			Price:            itemPrice,
		}
		items = append(items, itemModel)
	}

	err = r.DB.Create(&items).Error
	if err != nil {
		logrus.WithError(err).Error("Failed to save items to the database")
		return "", err
	}

	return receiptID, nil
}

// GetPoints returns the points for a receipt
func (r *ReceiptDB) GetPoints(id string) (int, error) {
	// Get the points for the receipt from the database
	return 0, nil
}
