package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onumahkalusamuel/hospital-app/config"
	"github.com/onumahkalusamuel/hospital-app/internal/api"
	"github.com/onumahkalusamuel/hospital-app/internal/api/delivery"
	"github.com/onumahkalusamuel/hospital-app/internal/api/invoice"
	"github.com/onumahkalusamuel/hospital-app/internal/api/patient"
	"github.com/onumahkalusamuel/hospital-app/internal/api/payment"
	"github.com/onumahkalusamuel/hospital-app/internal/api/staff"
	"github.com/onumahkalusamuel/hospital-app/internal/db"
	"github.com/onumahkalusamuel/hospital-app/internal/middleware"
	"github.com/onumahkalusamuel/hospital-app/web"
)

func main() {
	// get router
	e := echo.New()
	e.Use(echomiddleware.CORS())
	e.Static("files/", "files")

	// setup the database
	db.Init()
	db.InitialData()
	// set route for web files
	web.RegisterHandlers(e)

	unauthApi := e.Group("/api", middleware.ActivationAndInitialSetup())
	unauthApi.POST("/login", api.Login)
	unauthApi.POST("/activate", api.Activate)
	unauthApi.POST("/create-admin", api.CreateAdmin)
	unauthApi.POST("/hospital-setup", api.HospitalSetup)
	unauthApi.GET("/hospital-details", api.HospitalDetails)
	unauthApi.GET("/installation-code", api.InstallationCode)

	unauthApi.POST("/get-activation-code", api.GetActivationCode) // move endpoint to server when done

	authApi := e.Group("/api", middleware.ActivationAndInitialSetup(), echojwt.JWT([]byte(config.JwtSecret)), middleware.AppendJwtClaims())
	// profile actions
	authApi.GET("/profile", staff.ReadProfile)
	authApi.PUT("/profile", staff.UpdatePassword)
	// staff crud
	authApi.GET("/staff", staff.ReadAll)
	authApi.POST("/staff", staff.Create)
	authApi.GET("/staff/:id", staff.ReadOne)
	authApi.PUT("/staff/:id", staff.Update)
	authApi.DELETE("/staff/:id", staff.Delete)

	// patients crud and actions
	authApi.GET("/patients", patient.ReadAll)
	authApi.POST("/patients", patient.Create)
	authApi.GET("/patients/:id", patient.ReadOne)
	authApi.PUT("/patients/:id", patient.Update)
	authApi.DELETE("/patients/:id", patient.Delete)
	// next of kin
	authApi.POST("/patients/:patient_id/next-of-kin", patient.PatientNextOfKin)
	// patient invoices
	authApi.GET("/patients/:patient_id/invoices", patient.PatientInvoice)
	// deliveries crud and actions
	authApi.GET("/deliveries", delivery.ReadAll)
	authApi.POST("/deliveries", delivery.Create)
	authApi.GET("/deliveries/:id", delivery.ReadOne)
	authApi.PUT("/deliveries/:id", delivery.Update)
	authApi.DELETE("/deliveries/:id", delivery.Delete)
	// invoices crud and actions
	authApi.GET("/invoices", invoice.ReadAll)
	authApi.POST("/invoices", invoice.Create)
	authApi.GET("/invoices/:id", invoice.ReadOne)
	authApi.PUT("/invoices/:id", invoice.Update)
	authApi.DELETE("/invoices/:id", invoice.Delete)
	// payments crud and actions
	authApi.POST("/invoices/:invoice_id/payments", payment.Create)
	authApi.DELETE("/invoices/:invoice_id/payments/:id", payment.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
