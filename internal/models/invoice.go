package models

import (
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/pkg"
	"gorm.io/gorm"
)

type Invoice struct {
	BaseModel
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
	StaffID   string           `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID string           `gorm:"not null;references:patients(id)" json:"patient_id"`
	Details   []InvoiceDetails `gorm:"types:bytes;serializer:gob" json:"details"`
	Amount    uint             `gorm:"not null" json:"amount"`
	Balance   uint             `gorm:"default:0" json:"balance"`
	Completed uint             `gorm:"default:0" json:"completed"`
	Payments  []*Payment       `gorm:"constraint:onDelete:CASCADE" json:"payments"`
	Patient   *Patient         `json:"patient"`
}

type InvoiceDetails struct {
	Description string `json:"description"`
	Qty         uint   `json:"qty"`
	Unit        string `json:"unit"`
	Price       uint   `json:"price"`
	Amount      uint   `json:"amount"`
}

func (m *Invoice) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *Invoice) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Invoice) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Invoice) ReadAll() (bool, []Invoice) {
	var Invoices []Invoice
	if result := config.DB.Find(&Invoices, &m); result.Error != nil {
		return false, Invoices
	}
	return true, Invoices
}

func (cg *Invoice) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var invoices []*Invoice

	config.DB.Scopes(Paginate(invoices, &pagination, config.DB)).Find(&invoices)
	pagination.Rows = invoices

	return &pagination, nil
}
