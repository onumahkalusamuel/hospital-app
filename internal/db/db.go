package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	var err error

	// load env file if on localhost
	_, err = os.ReadFile(".env")
	if err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	if os.Getenv("ENV") == "dev" {
		config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	} else {

		// dsn := fmt.Sprintf(
		// 	"host=%v user=%v password=%v dbname=%v port=%v sslmode=enable TimeZone=Africa/Lagos",
		// 	os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
		// )
		dsn := os.Getenv("DATABASE_URL")

		config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

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
