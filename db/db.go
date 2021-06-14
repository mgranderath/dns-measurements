package db

import (
	"github.com/mgranderath/dns-measurements/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	dbConnect, err := gorm.Open(sqlite.Open("measurements.db?cache=shared&mode=rwc&_journal_mode=WAL"), &gorm.Config{

	})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbConnect

	// Migrate the schema
	err = db.AutoMigrate(&model.DNSMeasurement{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Traceroute{})
	if err != nil {
		log.Fatal(err)
	}
}

func AddMeasurement(measurement model.DNSMeasurement) {
	db.Create(&measurement)
}

func AddTraceroute(traceroute model.Traceroute) {
	db.Create(&traceroute)
}

func Close() {
	rawDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	rawDB.Close()
}