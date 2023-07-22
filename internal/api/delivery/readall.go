package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

func ReadAll(c echo.Context) error {
	// var deliveries []models.Delivery
	// config.DB.Model(&models.Delivery{}).Preload("Patient").Find(&deliveries)
	// return c.JSON(http.StatusOK, deliveries)

	handle := models.Delivery{}
	list, err := handle.List(pkg.GrabFromContext(pkg.Pagination{}, c), []string{"Patient"})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, list)
}
