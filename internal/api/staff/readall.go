package staff

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

func ReadAll(c echo.Context) error {
	p := models.Staff{}
	pag, err := p.List(pkg.GrabFromContext(pkg.Pagination{}, c), []string{})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pag)
}
