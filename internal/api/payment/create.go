package payment

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/models"
)

func Create(c echo.Context) error {

	var payment models.Payment
	if err := c.Bind(&payment); err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	if payment.AmountPaid == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "please enter all required fields."})
	}

	invoiceId := c.Param("invoice_id")

	// fetch the invoice
	var invoice models.Invoice
	invoice.ID = invoiceId
	if err := config.DB.Model(&models.Invoice{}).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "invoice not found." + err.Error()})
	}

	// check if amount in balance is above or equal to paid amount
	if invoice.Balance < payment.AmountPaid {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "paid amount is above the remaining balance of " + fmt.Sprint(invoice.Balance)})
	}

	// create the payment
	payment.InvoiceID = invoiceId
	payment.StaffID = c.Get("ID").(string)
	payment.Balance = invoice.Balance - payment.AmountPaid

	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "error creating record: " + err.Error()})
	}

	// save the new balance and completion status
	invoice.Balance = payment.Balance
	if invoice.Balance == 0 {
		invoice.Completed = 1
	}

	if err := config.DB.Save(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "error updating balance: " + err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "record created successfully"})
}
