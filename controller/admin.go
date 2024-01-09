package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
	}
	InitTemp.Temp.ExecuteTemplate(w, "Admin", nil)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
	}
	InitTemp.Temp.ExecuteTemplate(w, "error", nil)
}
