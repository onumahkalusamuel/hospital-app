package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

func PatientHistory(c echo.Context) error {

	p := models.PatientHistory{PatientID: c.Param("patient_id")}
	pag, err := p.List(pkg.GrabFromContext(pkg.Pagination{}, c), &models.PatientHistory{PatientID: c.Param("patient_id")})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pag)

	// var patientHistory []models.PatientHistory
	// db := config.DB.Model(&models.PatientHistory{}).Preload("Patient")

	// db.Order("date ASC")

	// db.Find(&patientHistory, &models.PatientHistory{PatientID: c.Param("patient_id")})

	// if len(patientHistory) < 1 {
	// 	return c.JSON(http.StatusNotFound, echo.Map{"message": "record not found"})
	// }

	// return c.JSON(http.StatusOK, patientHistory)
}
