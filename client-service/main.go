package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Hello, World!")
    app := gin.Default()
    route := app.Group("/client-service")

    route.GET("/health", HealthCheck)

    route.POST("/appointments/:category", ReserveAppointment)
    route.GET("/appointments/:patient_id", GetPatientAppointment)
    route.GET("/records/:patient_id", GetPatientRecords)
    route.POST("/records/:patient_id", UpdatePatientRecord)
    
    app.Run(":8080")
}
