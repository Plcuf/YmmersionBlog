package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"net/http"
	"fmt"
)

func Accueil(w http.ResponseWriter, r *http.Request) {	
	var err error
	InitStruct.Articles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		return
	}
	InitStruct.Articles.Jeux[0].Name = "caca"
	InitStruct.EditJSON(InitStruct.Articles)
	InitTemp.Temp.ExecuteTemplate(w, "Accueil", nil)
}



func Detail(w http.ResponseWriter, r *http.Request) {

	InitTemp.Temp.ExecuteTemplate(w, "Detail", nil)
}

func Categorie(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Categorie", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Add", nil)
}

func InitAdd(w http.ResponseWriter, r *http.Request) {
	// InitStruct.Articles.Titre := r.FormValue("titre")
	// InitStruct.Articles.Description := r.FormValue("Description")
	// InitStruct.Articles.Author := r.FormValue("Author")
	// InitStruct.Articles.Image := r.FormValue("Image")
	// InitStruct.Articles.DateCreated := r.FormValue("DateCreated")
	// InitStruct.lstArticles = append(lstArticles, InitStruct.Articles)

}
