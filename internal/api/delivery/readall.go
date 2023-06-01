package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadAll(c echo.Context) error {
	d := &models.Delivery{}
	_, deliveries := d.ReadAll()
	return c.JSON(http.StatusOK, deliveries)
}
