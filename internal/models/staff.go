package models

import (
	"github.com/onumahkalusamuel/hospital-app/config"
	"gorm.io/gorm"
)

type Staff struct {
	BaseModel
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Firstname  string         `gorm:"not null" json:"firstname"`
	Middlename string         `gorm:"default:null" json:"middlename"`
	Lastname   string         `gorm:"not null" json:"lastname"`
	Phone      string         `gorm:"default:null" json:"phone"`
	Email      string         `gorm:"default:null" json:"email"`
	Role       uint           `gorm:"default:3" json:"role"`
	Username   string         `gorm:"not null;unique" json:"username"`
	Password   string         `gorm:"not null" json:"-"`
}

func (m *Staff) Create() error {
	return config.DB.Create(&m).Error
}

func (m *Staff) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *Staff) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

func (m *Staff) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Staff) Read() error {
	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Staff) ReadAll() (bool, []Staff) {
	var Staffs []Staff
	if result := config.DB.Find(&Staffs, &m); result.Error != nil {
		return false, Staffs
	}
	return true, Staffs
}
