package main

type AppointmentResponse struct {
	ID        int    `json:"id"`
	DoctorID  int    `json:"doctor_id"`
	PatientID int    `json:"patient_id"`
	Hospital  string `json:"hospital"`
	Fee       int    `json:"fee"`
	Room      string `json:"room"`
	Confirmed bool   `json:"confirmed"`
}

type PatientRecord struct {
	PatientID  int                 `json:"patient_id"`
	DoctorID   int                 `json:"doctor_id"`
	Symptoms   map[string][]string `json:"symptoms"`
	Treatments map[string][]string `json:"treatments"`
}