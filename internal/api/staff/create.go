package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

type CreateStaffRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       uint   `json:"role"`
}

func Create(c echo.Context) error {

	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var req CreateStaffRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if req.Firstname == "" || req.Lastname == "" || req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	// check username
	var inuse = []models.Staff{}
	config.DB.Unscoped().Where("username='" + req.Username + "'").Find(&inuse)

	if len(inuse) > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "username already in use"})
	}

	staff := &models.Staff{
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Middlename: req.Middlename,
		Role:       req.Role,
		Username:   req.Username,
		Password:   pkg.HashPassword(req.Password),
	}

	if err := staff.Create(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating account: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": staff.ID})
}
