package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

type UpdateStaffRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       uint   `json:"role"`
}

func Update(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var req UpdateStaffRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	staff := &models.Staff{}
	staff.ID = c.Param("id")
	staff.Read()

	if err := staff.Read(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if req.Firstname != staff.Firstname {
		staff.UpdateSingle("firstname", req.Firstname)
	}

	if req.Lastname != staff.Lastname {
		staff.UpdateSingle("lastname", req.Lastname)
	}

	if req.Middlename != staff.Middlename {
		staff.UpdateSingle("middlename", req.Middlename)
	}

	if req.Role != staff.Role {
		staff.UpdateSingle("role", req.Role)
	}

	if req.Password != "" {
		staff.UpdateSingle("password", pkg.HashPassword(req.Password))
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record updated successfully."})
}
