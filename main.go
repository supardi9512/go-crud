package main

import (
	"go-crud/controllers/patientcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", patientcontroller.Index)
	http.HandleFunc("/patient", patientcontroller.Index)
	http.HandleFunc("/patient/index", patientcontroller.Index)
	http.HandleFunc("/patient/add", patientcontroller.Add)
	http.HandleFunc("/patient/edit", patientcontroller.Edit)

	http.ListenAndServe(":3000", nil)
}
