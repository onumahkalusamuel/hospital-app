package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Update(c echo.Context) error {
	if c.Get("Role").(float64) > config.ADMIN_ROLE {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized access"})
	}

	var invoice models.Invoice
	var onRecord models.Invoice

	if err := c.Bind(&invoice); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	invoice.Amount = 0
	for x, details := range invoice.Details {
		invoice.Details[x].Amount = details.Qty * details.Price
		invoice.Amount += invoice.Details[x].Amount
	}

	onRecord.ID = c.Param("id")
	config.DB.Model(&models.Invoice{}).Preload("Payments").First(&onRecord)

	if len(onRecord.Payments) > 0 {
		var paymentSum uint

		paymentSum = 0
		for _, payment := range onRecord.Payments {
			paymentSum += payment.AmountPaid
		}

		if paymentSum > invoice.Amount {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "payment records exceed the current invoice total. Please review."})
		}
	}

	invoice.ID = c.Param("id")

	config.DB.Updates(invoice)

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record updated successfully."})
}
