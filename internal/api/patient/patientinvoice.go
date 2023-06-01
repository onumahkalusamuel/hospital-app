package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientInvoice(c echo.Context) error {
	patient := &models.Patient{}
	patient.ID = c.Param("patient_id")
	config.DB.Model(&models.Patient{}).Preload("Invoices").First(&patient)

	if patient.CardNo == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, patient)
}
