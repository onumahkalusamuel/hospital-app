package invoice

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Create(c echo.Context) error {

	var invoice models.Invoice
	if err := c.Bind(&invoice); err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	fmt.Println(invoice)
	if invoice.PatientID == "" || len(invoice.Details) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "please enter all required fields."})
	}

	invoice.Amount = 0
	for x, details := range invoice.Details {
		invoice.Details[x].Amount = details.Qty * details.Price
		invoice.Amount += invoice.Details[x].Amount
	}

	invoice.StaffID = c.Get("ID").(string)
	invoice.Balance = invoice.Amount

	if err := config.DB.Create(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "error creating record: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record created successfully"})
}
