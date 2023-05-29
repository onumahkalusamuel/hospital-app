package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Update(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var delivery models.Delivery

	if err := c.Bind(&delivery); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	delivery.ID = c.Param("id")

	config.DB.Updates(delivery)

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record updated successfully."})
}
