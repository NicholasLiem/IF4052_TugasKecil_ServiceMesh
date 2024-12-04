package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	grandOakURL := "http://grand-oak-service/grand-oak/appointments/" + category
	pineValleyURL := "http://pine-valley-service/pine-valley/appointments/" + category

	requestBody, err := json.Marshal(appointmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Failed to encode request"})
		return
	}

	var message string
	var pineValleyResponse, grandOakResponse *AppointmentResponse
	successfulServices := 0

	resp, err = client.Post(pineValleyURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil || resp.StatusCode == 500 {
		fmt.Println("Error communicating with Pine Valley or received 500")
		pineValleyResponse = nil
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&pineValleyResponse); err != nil {
			pineValleyResponse = nil
		} else {
			successfulServices++
		}
	}
	fmt.Printf("Response from Pine Valley: %+v\n", pineValleyResponse)

	resp, err = client.Post(grandOakURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil || resp.StatusCode == 500 {
		fmt.Println("Error communicating with Grand Oak or received 500")
		grandOakResponse = nil
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&grandOakResponse); err != nil {
			grandOakResponse = nil
		} else {
			successfulServices++
		}
	}
	fmt.Printf("Response from Grand Oak: %+v\n", grandOakResponse)

	if successfulServices == 2 {
		message = "Appointment reserved successfully in both services."
	} else if successfulServices == 1 {
		message = "Appointment reserved successfully in one service."
	} else {
		message = "Failed to reserve appointments in both services."
	}

	ctx.JSON(200, gin.H{
		"message":              message,
		"pine_valley_response": pineValleyResponse,
		"grand_oak_response":   grandOakResponse,
	})
}

func GetPatientAppointment(ctx *gin.Context){
	patientId := ctx.Param("patient_id")

	if patientId == ""{
		ctx.JSON(400, gin.H{"Error": "Invalid Patient ID"})
		return
	}

	client := CreateHTTPClient()
	
	grandOakURL := "http://localhost:8082/grand-oak/appointments/" + patientId
	pineValleyURL := "http://localhost:8083/pine-valley/appointments/" + patientId

	var (
		message string
		pineValleyResponse, grandOakResponse *[]AppointmentResponse
	)

	successfulServices := 0

	resp, err := client.Get(pineValleyURL)

	if err != nil || resp.StatusCode == 500 {
		fmt.Println("Error communicating with Pine Valley or received 500")
		pineValleyResponse = nil
	} else if resp.StatusCode == 404{
		fmt.Println("Patient's appointment not found")
		pineValleyResponse = nil
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&pineValleyResponse); err != nil {
			pineValleyResponse = nil
		} else {
			successfulServices++
		}
	}
	fmt.Printf("Response from Pine Valley: %+v\n", pineValleyResponse)


	resp, err = client.Get(grandOakURL)
	if err != nil || resp.StatusCode == 500 {
		fmt.Println("Error communicating with Grand Oak or received 500")
		grandOakResponse = nil
	} else if resp.StatusCode == 404{
		fmt.Println("Patient's Appointment not found")
		grandOakResponse = nil
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&grandOakResponse); err != nil {
			grandOakResponse = nil
		} else {
			successfulServices++
		}
	}
	fmt.Printf("Response from Grand Oak: %+v\n", pineValleyResponse)

	if successfulServices == 2{
		message = "Success getting patient's appointment data from both services"
	}else if successfulServices == 1{
		message = "Patient's appointment succesfully pulled from 1 service only"
	}else{
		message = "Failed to get patient's appointment data from both services"
	}

	ctx.JSON(200, gin.H{
		"message":              message,
		"pine_valley_response": pineValleyResponse,
		"grand_oak_response":   grandOakResponse,
	})
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