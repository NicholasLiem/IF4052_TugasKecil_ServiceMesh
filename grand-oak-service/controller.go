package main

import (
	"fmt"
	"strings"
)

var MockAppointments = getMockAppointments()
var MockDoctors = getMockDoctors()
var MockPatients = getMockPatients()
var MockPatientRecords = getMockPatientRecords()

func CreateNewAppointment(doctorID, patientID int, hospital, category string) (*Appointment, error) {
	if !contain(MockCategories, category) {
		return nil, fmt.Errorf("invalid category %s", category)
	}

	doctor := FindDoctorByID(doctorID)
	patient := FindPatientByID(patientID)

	if patient == nil {
		return nil, fmt.Errorf("patient not found for ID %d", patientID)
	}

	if doctor == nil {
		return nil, fmt.Errorf("doctor not found for ID %d", doctorID)
	}

	if !strings.EqualFold(doctor.Category, category) {
		return nil, fmt.Errorf("doctor is not available for category %s", category)
	}

	appointment := Appointment{
		ID:        len(MockAppointments) + 1,
		DoctorID:  doctor.ID,
		PatientID: patient.ID,
		Hospital:  hospital,
		Fee:       doctor.Fee,
		Confirmed: false,
	}
	MockAppointments = append(MockAppointments, appointment)
	return &appointment, nil
}

func getPatientRecord(patientID int) *PatientRecord {
	for _, record := range MockPatientRecords {
		if record.PatientID == patientID {
			return &record
		}
	}
	return nil
}

func updatePatientRecord(patientID int, symptoms, treatments []string) (*PatientRecord, error) {
	record := getPatientRecord(patientID)
	if record == nil {
		return nil, fmt.Errorf("record not found for patient ID %d", patientID)
	}

	record.UpdateSymptoms(symptoms)
	record.UpdateTreatments(treatments)

	return record, nil
}

func getAppointmentsByPatientID(patientID int) []Appointment {
	var appointments []Appointment
	for _, appointment := range MockAppointments {
		if appointment.PatientID == patientID {
			appointments = append(appointments, appointment)
		}
	}
	return appointments
}

func FindDoctorByID(id int) *Doctor {
	for _, doctor := range MockDoctors {
		if doctor.ID == id {
			return &doctor
		}
	}
	return nil
}

func FindPatientByID(id int) *Patient {
	for _, patient := range MockPatients {
		if patient.ID == id {
			return &patient
		}
	}
	return nil
}

func contain(arr []string, str string) bool {
	for _, s := range arr {
		if strings.EqualFold(s, str) {
			return true
		}
	}
	return false
}
