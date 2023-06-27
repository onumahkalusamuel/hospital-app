package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientHistoryUpdate(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var patientHistory models.PatientHistory

	if err := c.Bind(&patientHistory); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	patientHistory.ID = c.Param("hid")
	patientHistory.PatientID = c.Param("patient_id")

	config.DB.Updates(patientHistory)

	return c.JSON(http.StatusOK, echo.Map{"message": "record updated successfully."})
}
