package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	var (
		resp *http.Response
	)

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

	client := CreateHTTPClient()
	// http://grand-oak-service pas deploy ini buat local aja
	// http://pine-valley-service
	grandOakURL := "http://localhost:8080/grand-oak/appointments/" + category
	pineValleyURL := "http://localhost:8080/pine-valley/appointments/" + category

	requestBody, err := json.Marshal(appointmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Failed to encode request"})
		return
	}


	resp, err = client.Post(pineValleyURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Failed to communicate with Pine Valley"})
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response from Pine Valley: %s\n", body)

	// Hit Grand Oak
	resp, err = client.Post(grandOakURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Failed to communicate with Grand Oak"})
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("Response from Grand Oak: %s\n", body)

	ctx.JSON(200, gin.H{"message": "Appointment reserved in both services", "body": string(body)})
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

func HealthCheck(ctx *gin.Context){
	ctx.JSON(200, gin.H{"Status": "OK"})
}