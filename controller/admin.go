package controller

import (
	InitTemp "Ymmersion2/temps"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Admin", nil)
}

func Error(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Error", nil)
}