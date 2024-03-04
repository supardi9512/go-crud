package models

import (
	"database/sql"
	"go-crud/config"
	"go-crud/entities"
	"time"
)

type PatientModel struct {
	conn *sql.DB
}

func NewPatientModel() *PatientModel {
	conn, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &PatientModel{
		conn: conn,
	}
}

func (p *PatientModel) FindAll() ([]entities.Patient, error) {
	rows, err := p.conn.Query("SELECT * FROM patients")

	if err != nil {
		return []entities.Patient{}, err
	}

	defer rows.Close()

	var patientData []entities.Patient

	for rows.Next() {
		var patient entities.Patient
		rows.Scan(&patient.Id, &patient.Name, &patient.Nik, &patient.Gender, &patient.PlaceOfBirth, &patient.DateOfBirth, &patient.Address, &patient.PhoneNumber)

		if patient.Gender == "1" {
			patient.Gender = "Male"
		} else if patient.Gender == "2" {
			patient.Gender = "Female"
		}

		// 2006-01-02 => yyyy-mm-dd
		date_of_birth, _ := time.Parse("2006-01-02", patient.DateOfBirth)

		// 02-01-2006 => dd-mm-yyyy
		patient.DateOfBirth = date_of_birth.Format("02-01-2006")

		patientData = append(patientData, patient)
	}

	return patientData, nil
}

func (p *PatientModel) Create(patient entities.Patient) bool {
	result, err := p.conn.Exec("INSERT INTO patients (name, nik, gender, place_of_birth, date_of_birth, address, phone_number) VALUES(?,?,?,?,?,?,?)", patient.Name, patient.Nik, patient.Gender, patient.PlaceOfBirth, patient.DateOfBirth, patient.Address, patient.PhoneNumber)

	if err != nil {
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
