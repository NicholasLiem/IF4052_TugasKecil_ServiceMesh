package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentRequest struct {
	DoctorID int 	`json:"doctor_id"`
	PatientID int	`json:"patient_id"`
}

type RecordRequest struct {
	Symptoms   []string `json:"symptoms"`
	Treatments []string `json:"treatments"`
}

func ReserveAppointment(ctx *gin.Context){
	category := ctx.Param("category")

	if category == ""{
		ctx.JSON(400, gin.H{"Error": "Invalid Category!"})
		return
	}

	var appointmentRequest AppointmentRequest

	if err := ctx.ShouldBindJSON(&appointmentRequest); err != nil{
		ctx.JSON(400, gin.H{"Error": "Invalid request body"})
		return
	}

	// Hit ke pine valley sama grand oak

}

func GetPatientAppointment(ctx *gin.Context){
	fmt.Println("MASUK KE SINI")
	patientId := ctx.Param("patient_id")

	if patientId == ""{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}

	id, err := strconv.Atoi(patientId)
	
	if err != nil{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}

	fmt.Println("INI ID", id)

	// Ini hit ke pine valley sama ke grand oak
}
 
func GetPatientRecords(ctx *gin.Context){
	patientId := ctx.Param("patient_id")

	if patientId == ""{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}

	id, err := strconv.Atoi(patientId)

	if err != nil{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}

	fmt.Println("INI id", id)

	// Ini hit ke pine valley sama ke grand oak
}

func UpdatePatientRecord(ctx *gin.Context){
	patientId := ctx.Param("patient_id")

	if patientId == ""{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}
	var recordRequest RecordRequest
	if err := ctx.ShouldBindJSON(&recordRequest); err != nil{
		ctx.JSON(400, gin.H{"Error": "Invalid Request Body"})
		return
	}

	// Ini hit ke pine valley sama ke grand oak
	fmt.Println("INI RECORD REQUEST", recordRequest)
}