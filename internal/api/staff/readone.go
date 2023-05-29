package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadOne(c echo.Context) error {
	staff := &models.Staff{}
	staff.ID = c.Param("id")
	staff.Read()
	if staff.Firstname == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"success": false, "message": "Staff record not found"})
	}

	staff.Password = ""
	return c.JSON(http.StatusOK, echo.Map{"success": true, "data": staff})
}
