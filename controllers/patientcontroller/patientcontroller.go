package patientcontroller

import (
	"go-crud/entities"
	"go-crud/models"
	"html/template"
	"net/http"
)

var patientModel = models.NewPatientModel()

func Index(response http.ResponseWriter, request *http.Request) {

	patient, _ := patientModel.FindAll()

	data := map[string]interface{}{
		"patient": patient,
	}

	temp, err := template.ParseFiles("views/patient/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient

		patient.Name = request.Form.Get("name")
		patient.Nik = request.Form.Get("nik")
		patient.Gender = request.Form.Get("gender")
		patient.PlaceOfBirth = request.Form.Get("place_of_birth")
		patient.DateOfBirth = request.Form.Get("date_of_birth")
		patient.Address = request.Form.Get("address")
		patient.PhoneNumber = request.Form.Get("phone_number")

		patientModel.Create(patient)

		data := map[string]interface{}{
			"message": "Patient data has been added successfully",
		}

		temp, _ := template.ParseFiles("views/patient/add.html")

		temp.Execute(response, data)
	}
}
