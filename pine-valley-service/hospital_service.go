package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentRequest struct {
	DoctorID  int `json:"doctor_id"`
	PatientID int `json:"patient_id"`
}

type RecordRequest struct {
	Symptoms   []string `json:"symptoms"`
	Treatments []string `json:"treatments"`
}

var MockCategories = getMockCategories()
var Hospital = "Pine Valley Hospital"

func ReserveAppointment(c *gin.Context) {
	category := c.Param("category")

	if category == "" {
		c.JSON(400, gin.H{"error": "Invalid category"})
		return
	}

	var appointmentRequest AppointmentRequest
	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	appointment, err := CreateNewAppointment(
		appointmentRequest.DoctorID,
		appointmentRequest.PatientID,
		Hospital, category)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointment)
}

func GetAppointments(c *gin.Context) {
	patientIDStr := c.Param("patient_id")

	if patientIDStr == "" {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	appointments := getAppointmentsByPatientID(patientID)

	if len(appointments) == 0 {
		c.JSON(404, gin.H{"error": "No appointments found"})
		return
	}

	c.JSON(200, appointments)
}

func GetPatientRecords(c *gin.Context) {
	patientIDStr := c.Param("patient_id")

	if patientIDStr == "" {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	record := getPatientRecord(patientID)

	if record == nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, record)
}

func UpdatePatientRecord(c *gin.Context) {
	patientIDStr := c.Param("patient_id")

	if patientIDStr == "" {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	var recordRequest RecordRequest
	if err := c.ShouldBindJSON(&recordRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	record, err := updatePatientRecord(patientID, recordRequest.Symptoms, recordRequest.Treatments)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, record)
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}