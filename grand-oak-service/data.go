package main

func getMockAppointments() []Appointment {
	return []Appointment{
		{ID: 201, DoctorID: 301, PatientID: 401, Hospital: "Grand Oak Hospital", Fee: 200, Room: "Room 101", Confirmed: true},
		{ID: 202, DoctorID: 302, PatientID: 402, Hospital: "Grand Oak Hospital", Fee: 300, Room: "Room 202", Confirmed: false},
	}
}

func getMockDoctors() []Doctor {
	return []Doctor{
		{ID: 301, Name: "Dr. Amanda Carter", Hospital: "Grand Oak Hospital", Category: "Pediatrics", Availability: "9:00 AM - 5:00 PM", Fee: 200, Room: "Room 101"},
		{ID: 302, Name: "Dr. Brian Taylor", Hospital: "Grand Oak Hospital", Category: "Orthopedics", Availability: "10:00 AM - 6:00 PM", Fee: 300, Room: "Room 202"},
		{ID: 303, Name: "Dr. Christine Adams", Hospital: "Grand Oak Hospital", Category: "Dermatology", Availability: "8:00 AM - 4:00 PM", Fee: 250, Room: "Room 303"},
		{ID: 304, Name: "Dr. David Wilson", Hospital: "Grand Oak Hospital", Category: "Neurology", Availability: "11:00 AM - 7:00 PM", Fee: 400, Room: "Room 404"},
	}
}

func getMockCategories() []string {
	return []string{"General", "Surgery", "Pediatrics", "Orthopedics", "Dermatology", "Neurology"}
}

func getMockPatients() []Patient {
	return []Patient{
		{ID: 401, Name: "Ethan Green", Dob: "05-06-1985", SSN: "456-78-9012", Address: "123 Oak Street"},
		{ID: 402, Name: "Sophia Brown", Dob: "12-08-1990", SSN: "654-32-1098", Address: "456 Oak Avenue"},
	}
}

func getMockPatientRecords() []PatientRecord {
	return []PatientRecord{
		{PatientID: 401, DoctorID: 301, Symptoms: map[string][]string{"03-11-2024": {"Fever", "Cold"}}, Treatments: map[string][]string{"03-11-2024": {"Paracetamol"}}},
		{PatientID: 402, DoctorID: 302, Symptoms: map[string][]string{"05-11-2024": {"Back Pain"}}, Treatments: map[string][]string{"05-11-2024": {"Physiotherapy"}}},
	}
}
