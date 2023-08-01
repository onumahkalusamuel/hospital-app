package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

func ReadAll(c echo.Context) error {

	p := models.Invoice{}
	if patient_id := c.QueryParam("patient_id"); patient_id != "" {
		p.PatientID = patient_id
	}

	pag, err := p.List(pkg.GrabFromContext(pkg.Pagination{}, c), &p)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pag)
}
