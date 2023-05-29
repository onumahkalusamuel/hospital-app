package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadOne(c echo.Context) error {
	delivery := &models.Delivery{}
	delivery.ID = c.Param("id")
	delivery.Read()
	if delivery.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"success": false, "message": "record not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "data": delivery})
}
