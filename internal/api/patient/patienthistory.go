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
}
