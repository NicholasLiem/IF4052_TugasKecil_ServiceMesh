package main

func getMockAppointments() []Appointment {
	return []Appointment{
		{ID: 101, DoctorID: 101, PatientID: 101, Hospital: "Pine Valley Hospital", Fee: 100, Confirmed: true},
		{ID: 102, DoctorID: 102, PatientID: 102, Hospital: "Pine Valley Hospital", Fee: 150, Confirmed: false},
	}
}

func getMockDoctors() []Doctor {
	return []Doctor{
		{ID: 101, Name: "Dr. John Doe", Hospital: "Pine Valley Hospital", Category: "General", Availability: "9:00 AM - 5:00 PM", Fee: 100},
		{ID: 102, Name: "Dr. Jane Smith", Hospital: "Pine Valley Hospital", Category: "Surgery", Availability: "10:00 AM - 6:00 PM", Fee: 150},
		{ID: 103, Name: "Dr. Mike Brown", Hospital: "Pine Valley Hospital", Category: "Cardiology", Availability: "8:00 AM - 4:00 PM", Fee: 90},
		{ID: 104, Name: "Dr. Sarah White", Hospital: "Pine Valley Hospital", Category: "Gynaecology", Availability: "11:00 AM - 7:00 PM", Fee: 200},
	}
}

func getMockCategories() []string {
	return []string{"General", "Surgery", "Cardiology", "Gynaecology"}
}

func getMockPatients() []Patient {
	return []Patient{
		{ID: 101, Name: "Alice", Dob: "01-01-1990", SSN: "123-45-6789", Address: "123, Pine Valley"},
		{ID: 102, Name: "Bob", Dob: "02-02-1991", SSN: "987-65-4321", Address: "456, Pine Valley"},
	}
}

func getMockPatientRecord() []PatientRecord {
	return []PatientRecord{
		{PatientID: 101, DoctorID: 101, Symptoms: map[string][]string{"01-01-2024": {"Fever", "Cough"}, "02-01-2024": {"Headache"}}, Treatments: map[string][]string{"01-01-2024": {"Paracetamol", "Cough syrup"}}},
		{PatientID: 102, DoctorID: 102, Symptoms: map[string][]string{"01-01-2024": {"Stomach ache"}, "02-01-2024": {"Nausea"}}, Treatments: map[string][]string{"01-01-2024": {"Antacid"}}},
	}
}
