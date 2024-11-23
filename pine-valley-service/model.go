package main

import "time"

type Appointment struct {
	ID        int    `json:"id"`
	DoctorID  int    `json:"doctor_id"`
	PatientID int    `json:"patient_id"`
	Hospital  string `json:"hospital"`
	Fee       int    `json:"fee"`
	Confirmed bool   `json:"confirmed"`
}

type Doctor struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Hospital     string `json:"hospital"`
	Category     string `json:"category"`
	Availability string `json:"availability"`
	Fee          int    `json:"fee"`
}

type Patient struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Dob     string `json:"dob"`
	SSN     string `json:"ssn"`
	Address string `json:"address"`
}

type PatientRecord struct {
	PatientID  int                 `json:"patient_id"`
	DoctorID   int                 `json:"doctor_id"`
	Symptoms   map[string][]string `json:"symptoms"`
	Treatments map[string][]string `json:"treatments"`
}

func (pr *PatientRecord) UpdateTreatments(treatments []string) {
	date := time.Now().Format("02-01-2006")
	if pr.Treatments == nil {
		pr.Treatments = make(map[string][]string)
	}
	pr.Treatments[date] = treatments
}

func (pr *PatientRecord) UpdateSymptoms(symptoms []string) {
	date := time.Now().Format("02-01-2006")
	if pr.Symptoms == nil {
		pr.Symptoms = make(map[string][]string)
	}
	pr.Symptoms[date] = symptoms
}
