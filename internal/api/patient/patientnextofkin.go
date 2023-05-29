package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func PatientNextOfKin(c echo.Context) error {

	var nextofkin models.NextOfKin
	if err := c.Bind(&nextofkin); err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	if nextofkin.Firstname == "" || nextofkin.Relationship == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "please enter all required fields."})
	}

	patientId := c.Param("patient_id")
	// fetch the invoice
	var patient models.Patient
	patient.ID = patientId
	if err := config.DB.Model(&models.Patient{}).Preload("NextOfKin").First(&patient).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "patient not found." + err.Error()})
	}

	if patient.NextOfKin != nil && patient.NextOfKin.ID != "" {
		nextofkin.ID = patient.NextOfKin.ID
	}

	nextofkin.StaffID = c.Get("ID").(string)
	nextofkin.PatientID = patientId

	if err := config.DB.Save(&nextofkin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "next of kin updated successfully"})
}
