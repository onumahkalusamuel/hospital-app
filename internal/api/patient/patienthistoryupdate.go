package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/helpers"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientHistoryUpdate(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var patientHistory models.PatientHistory

	patientHistory.StaffID = c.Get("ID").(string)
	patientHistory.PatientID = c.Param("patient_id")
	patientHistory.Type = c.FormValue("type")
	patientHistory.Details = echo.Map{"note": c.FormValue("note")}

	if patientHistory.Type == "Admission" {
		patientHistory.Details["room_number"] = c.FormValue("room_number")
		patientHistory.Details["ward_number"] = c.FormValue("ward_number")
	}

	if patientHistory.Type == "Appointment" {
		patientHistory.Details["appointment_type"] = c.FormValue("appointment_type")
	}

	patientHistory.ID = c.Param("history_id")
	patientHistory.PatientID = c.Param("patient_id")

	// upload the file
	filename, err := helpers.SaveImage(c, "document", "images")

	if err == nil {
		patientHistory.Document = filename
	}
	config.DB.Updates(patientHistory)

	return c.JSON(http.StatusOK, echo.Map{"message": "record updated successfully."})
}
