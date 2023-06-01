package patient

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func ReadAll(c echo.Context) error {
	var patients []models.Patient

	db := config.DB.Model(&patients)

	if query := c.QueryParam("query"); query != "" {
		db.Where(
			"lastname LIKE '" + query + "%' OR " +
				"firstname LIKE '" + query + "%' OR " +
				"middlename LIKE '" + query + "%' OR " +
				"card_no LIKE '%" + query + "%' OR " +
				"phone LIKE '" + query + "%' OR " +
				"email LIKE '" + query + "%' OR " +
				"occupation LIKE '" + query + "%'")
	}

	db.Find(&patients)
	return c.JSON(http.StatusOK, patients)
}
