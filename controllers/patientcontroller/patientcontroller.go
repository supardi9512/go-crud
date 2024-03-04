package patientcontroller

import (
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/patient/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(response, nil)
}
