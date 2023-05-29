package models

import (
	"github.com/onumahkalusamuel/hospital-app/config"
)

type NextOfKin struct {
	BaseModel
	StaffID      string `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID    string `gorm:"not null;references:patients(id)" json:"patient_id"`
	Title        string `gorm:"default:null" json:"title"`
	Firstname    string `gorm:"not null" json:"firstname"`
	Middlename   string `gorm:"default:null" json:"middlename"`
	Lastname     string `gorm:"not null" json:"lastname"`
	Phone        string `gorm:"default:null" json:"phone"`
	Email        string `gorm:"default:null" json:"email"`
	Sex          string `gorm:"default:null" json:"sex"`
	Address      string `gorm:"default:null" json:"address"`
	Occupation   string `gorm:"default:null" json:"occupation"`
	Relationship string `gorm:"default:null" json:"relationship"`
}

func (m *NextOfKin) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *NextOfKin) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *NextOfKin) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *NextOfKin) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *NextOfKin) ReadAll() (bool, []NextOfKin) {
	var NextOfKins []NextOfKin
	if result := config.DB.Find(&NextOfKins, &m); result.Error != nil {
		return false, NextOfKins
	}
	return true, NextOfKins
}
