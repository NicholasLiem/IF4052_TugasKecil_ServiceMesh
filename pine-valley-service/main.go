package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	route := app.Group("/pine-valley")

	route.GET("/health", HealthCheck)

	route.POST("/appointments/:category", ReserveAppointment)

	route.GET("/appointments/:patient_id", GetAppointments)

	route.GET("/records/:patient_id", GetPatientRecords)

	route.POST("/records/:patient_id", UpdatePatientRecord)

	app.Run(":8080")
}
