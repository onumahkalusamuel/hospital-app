package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadProfile(c echo.Context) error {
	staff := &models.Staff{}
	staff.ID = c.Get("id").(string)
	staff.Read()
	if staff.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	}

	staff.Password = ""
	return c.JSON(http.StatusOK, staff)
}
