package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadAll(c echo.Context) error {
	var invoices []models.Invoice

	db := config.DB.Model(&models.Invoice{})

	if balance := c.QueryParam("balance"); balance == "true" {
		db.Where("balance > 0")
	} else if balance == "false" {
		db.Where("balance == 0")
	} else {
		db.Order("balance DESC, created_at ASC")
	}

	if patient_id := c.QueryParam("patient_id"); patient_id != "" {
		db.Where("patient_id = '" + patient_id + "'")
	}

	db.Preload("Patient").Find(&invoices)

	return c.JSON(http.StatusOK, invoices)
}
