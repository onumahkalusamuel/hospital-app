package patient

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/helpers"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientHistoryCreate(c echo.Context) error {

	var patientHistory models.PatientHistory

	now := time.Now()

	patientHistory.StaffID = c.Get("ID").(string)
	patientHistory.PatientID = c.Param("patient_id")
	patientHistory.Type = c.FormValue("type")
	patientHistory.Details = echo.Map{"note": c.FormValue("note")}
	patientHistory.Date = &now

	// upload the file
	filename, err := helpers.SaveImage(c, "document", "images")

	if err == nil {
		patientHistory.Document = filename
	}

	if err := config.DB.Create(&patientHistory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": patientHistory.ID})
}
