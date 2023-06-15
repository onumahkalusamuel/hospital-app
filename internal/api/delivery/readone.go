package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadOne(c echo.Context) error {
	var delivery models.Delivery
	delivery.ID = c.Param("id")
	config.DB.Model(&models.Delivery{}).Preload("Patient").First(&delivery)
	if delivery.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	return c.JSON(http.StatusOK, delivery)
}
