package config

import (
	"github.com/denisbrodbeck/machineid"
	"gorm.io/gorm"
)

const (
	UpdatePath = "./updates/"
	JwtSecret  = "secret"
	AppKey     = "hospital"
)

var (
	DB               *gorm.DB
	HashedID, err    = machineid.ProtectedID(AppKey)
	ActivationCode   = ""
	PaystackCurrency = "NGN"
	PaystackChannels = []string{"card", "bank", "ussd", "qr", "bank_transfer"}
)

// initial db settings
var (
	InitSettings = map[string]string{
		"installation_code": HashedID,
		"activation_code":   "",
		"last_card_no":      "1",
		"hospital_name":     "",
		"hospital_address":  "",
		"hospital_phone":    "",
		"hospital_email":    "",
		"hospital_logo":     "",
	}
)

const (
	ADMIN_ROLE = iota + 1
	DOCTOR_ROLE
	NURSE_ROLE
)

type PaystackMetaData struct {
	BlockGroupID     string
	UserID           string
	PaymentReference string
}

type BodyStructure map[string]string
