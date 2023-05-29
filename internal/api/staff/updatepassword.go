package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func UpdatePassword(c echo.Context) error {

	var req UpdatePasswordRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "bad request: " + err.Error()})
	}

	staff := &models.Staff{}
	staff.ID = c.Get("id").(string)
	staff.Read()

	if err := staff.Read(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "bad request: " + err.Error()})
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "enter old and new password."})
	}

	if pkg.CheckPassword(staff.Password, req.OldPassword) == false {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "wrong old password provided."})
	}

	staff.Password = pkg.HashPassword(req.NewPassword)
	config.DB.Model(&models.Staff{}).Updates(staff)

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "password changed successfully."})
}
