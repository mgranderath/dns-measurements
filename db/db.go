package db

import (
	"github.com/mgranderath/dns-measurements/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

var db *gorm.DB

func init() {
	dbConnect, err := gorm.Open(sqlite.Open("measurements.db"), &gorm.Config{

	})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbConnect

	// Migrate the schema
	err = db.AutoMigrate(&model.Server{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Measurement{})
	if err != nil {
		log.Fatal(err)
	}
}

func CreateServer(IP string, RTT *time.Duration) {
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&model.Server{IP: IP, RTT: RTT})
}

func AddMeasurement(measurement model.Measurement) {
	db.Create(&measurement)
}