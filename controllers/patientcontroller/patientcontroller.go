package patientcontroller

import (
	"go-crud/entities"
	"go-crud/libraries"
	"go-crud/models"
	"html/template"
	"net/http"
	"strconv"
)

var validation = libraries.NewValidation()
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

		var data = make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Patient data has been added successfully"
			patientModel.Create(patient)
		}

		temp, _ := template.ParseFiles("views/patient/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var patient entities.Patient
		patientModel.Find(id, &patient)

		data := map[string]interface{}{
			"patient": patient,
		}

		temp, err := template.ParseFiles("views/patient/edit.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var patient entities.Patient

		patient.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		patient.Name = request.Form.Get("name")
		patient.Nik = request.Form.Get("nik")
		patient.Gender = request.Form.Get("gender")
		patient.PlaceOfBirth = request.Form.Get("place_of_birth")
		patient.DateOfBirth = request.Form.Get("date_of_birth")
		patient.Address = request.Form.Get("address")
		patient.PhoneNumber = request.Form.Get("phone_number")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Patient data has been updated successfully"
			patientModel.Update(patient)
		}

		temp, _ := template.ParseFiles("views/patient/edit.html")
		temp.Execute(response, data)

	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	patientModel.Delete(id)

	http.Redirect(response, request, "/patient", http.StatusSeeOther)
}
