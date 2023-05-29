package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Create(c echo.Context) error {

	var delivery models.Delivery
	if err := c.Bind(&delivery); err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	fmt.Println(delivery)
	if delivery.BabySex == "" || delivery.PatientID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "please enter all required fields."})
	}

	delivery.StaffID = c.Get("ID").(string)

	if err := config.DB.Create(&delivery).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record created successfully"})
}
