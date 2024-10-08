package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadOne(c echo.Context) error {
	invoice := &models.Invoice{}
	invoice.ID = c.Param("id")
	config.DB.Model(&models.Invoice{}).Preload("Payments").Preload("Patient").First(&invoice)

	if invoice.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, invoice)
}
