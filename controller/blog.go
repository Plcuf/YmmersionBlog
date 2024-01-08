package controller 

import (
	InitTemp "Ymmersion2/temps"
	"net/http"
)

func Accueil (w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Accueil", nil)
}

func Detail (w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Detail", nil)
}

func Categorie (w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Categorie", nil)
}

func Add (w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Add", nil)
}