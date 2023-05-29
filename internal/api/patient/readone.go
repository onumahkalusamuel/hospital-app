package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadOne(c echo.Context) error {
	var patient models.Patient
	patient.ID = c.Param("id")
	config.DB.Model(&models.Patient{}).Preload("NextOfKin").First(&patient)
	if patient.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"success": false, "message": "account not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "data": patient})
}
