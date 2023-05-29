package models

import (
	"time"

	"github.com/onumahkalusamuel/hospital-app/config"
)

type Delivery struct {
	BaseModel
	StaffID          string     `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID        string     `gorm:"not null;references:patients(id)" json:"patient_id"`
	DeliveryDateTime *time.Time `gorm:"default:null" json:"delivery_date_time"`
	DeliveryMode     string     `gorm:"default:'vaginal'" json:"delivery_mode"`
	BabySex          string     `gorm:"not null" json:"baby_sex"`
	BabyWeight       uint       `gorm:"default:0" json:"baby_weight"`
	BabyWeightUnit   string     `gorm:"default:'kg'" json:"baby_weight_unit"`
	Condition        string     `gorm:"default:null" json:"condition"`
	Note             string     `gorm:"default:null" json:"note"`
	Patient          *Patient   `json:"patient"`
}

// Delete function
func (m *Delivery) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Delivery) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Delivery) ReadAll() (bool, []Delivery) {
	var Deliverys []Delivery
	if result := config.DB.Find(&Deliverys, &m); result.Error != nil {
		return false, Deliverys
	}
	return true, Deliverys
}
