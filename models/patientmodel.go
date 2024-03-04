package models

import (
	"database/sql"
	"go-crud/config"
	"go-crud/entities"
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

func (p PatientModel) Create(patient entities.Patient) bool {
	result, err := p.conn.Exec("INSERT INTO patients (name, nik, gender, place_of_birth, date_of_birth, address, phone_number) VALUES(?,?,?,?,?,?,?)", patient.Name, patient.Nik, patient.Gender, patient.PlaceOfBirth, patient.DateOfBirth, patient.Address, patient.PhoneNumber)

	if err != nil {
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
