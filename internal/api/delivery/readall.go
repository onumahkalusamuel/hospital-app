package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadAll(c echo.Context) error {
	var deliveries []models.Delivery
	config.DB.Model(&models.Delivery{}).Preload("Patient").Find(&deliveries)
	return c.JSON(http.StatusOK, deliveries)
}
