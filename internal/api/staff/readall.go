package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadAll(c echo.Context) error {
	d := &models.Staff{}
	_, staffs := d.ReadAll()
	return c.JSON(http.StatusOK, echo.Map{"success": true, "data": staffs})
}
