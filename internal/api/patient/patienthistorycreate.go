package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientHistoryCreate(c echo.Context) error {

	var patientHistory models.PatientHistory
	if err := c.Bind(&patientHistory); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request" + err.Error()})
	}

	if patientHistory.Details == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "please enter all required fields."})
	}

	patientHistory.StaffID = c.Get("ID").(string)
	patientHistory.PatientID = c.Param("patient_id")

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": patientHistory.ID})
}
