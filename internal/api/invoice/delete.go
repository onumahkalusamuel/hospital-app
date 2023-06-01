package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Delete(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	invoice := &models.Invoice{}
	invoice.ID = c.Param("id")
	config.DB.Model(&models.Invoice{}).Preload("Payments").First(&invoice)

	if invoice.PatientID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	config.DB.Delete(invoice)

	return c.JSON(http.StatusOK, echo.Map{"message": "record deleted successfully"})
}
