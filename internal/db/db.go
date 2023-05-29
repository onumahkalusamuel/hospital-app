package db

import (
	"log"

	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	var err error
	config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	config.DB.AutoMigrate(
		&models.Delivery{},
		&models.Invoice{},
		&models.NextOfKin{},
		&models.Patient{},
		&models.PatientHistory{},
		&models.Payment{},
		&models.Settings{},
		&models.Staff{},
	)
}

func InitialData() {
	for setting, value := range config.InitSettings {
		s := &models.Settings{Setting: setting}
		s.Read()
		if s.ID == "" {
			s = &models.Settings{Setting: setting, Value: value}
			s.Create()
		}
	}
}
