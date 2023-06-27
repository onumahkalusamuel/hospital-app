package models

import (
	"time"

	"github.com/onumahkalusamuel/hospital-app/config"
)

type PatientHistory struct {
	BaseModel
	StaffID   string                 `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID string                 `gorm:"not null;references:patients(id)"  json:"-"`
	Date      *time.Time             `gorm:"default:null" json:"date_time"`
	Type      string                 `gorm:"not null" json:"type"`
	Details   map[string]interface{} `gorm:"serializer:json" json:"details"`
	Patient   *Patient               `json:"patient"`
}

/*
Types:
payment (payment id, invoice id, payment amount, balance),
invoice (invoice id, amount, status),
note (subject, body),
admit (ward, room),
discharge (room),
history (note),
examination (note, image),
test-result (note, image),
diagnosis (note, image),
treatment (),
prescription (drug, dosage, frequency, timeline),
*/

func (m *PatientHistory) Create() error {
	return config.DB.Create(&m).Error
}

func (m *PatientHistory) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *PatientHistory) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

func (m *PatientHistory) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *PatientHistory) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *PatientHistory) ReadAll() (bool, []PatientHistory) {
	var PatientHistorys []PatientHistory
	if result := config.DB.Find(&PatientHistorys, &m); result.Error != nil {
		return false, PatientHistorys
	}
	return true, PatientHistorys
}

func (m *PatientHistory) UpdateLastPing() error {
	// .Format("2006-01-02")
	return config.DB.Model(&m).Update("last_ping", time.Now()).Error
}
