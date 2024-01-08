package controller

import (
	InitTemp "Ymmersion2/temps"
	InitStruct "Ymmersion2/backend"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
	}
	InitTemp.Temp.ExecuteTemplate(w, "Admin", nil)
}

func Error(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Error", nil)
}