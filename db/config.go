package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Connect connects to the database
func Connect() {
	var err error
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to the database")
	}
	logrus.Info("Connected to the database")

	// Migrate the schema
	logrus.Info("Migrating the schema")
	db.AutoMigrate(&Receipt{})
	db.AutoMigrate(&Item{})
	logrus.Info("Schema migrated")
}
