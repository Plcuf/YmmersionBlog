package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"fmt"
	"net/http"
	"os"
)

var err error

func Accueil(w http.ResponseWriter, r *http.Request) {
	var lstId []int
	var Recommandation []Article
	InitStruct.LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	lstId = InitStruct.NbAleatoire(InitStruct.LstArticles)
	for _, i := range lstId {
		Recommandation = append(Recommandation, LstArticles[i])
	}
	InitTemp.Temp.ExecuteTemplate(w, "Accueil", Recommandation)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryID := r.URL.Query().Get("id")
	for _, i := range InitStruct.LstArticles {
		if i.Id == queryID {
			InitStruct.Section = i
			break
		}
	}
	InitTemp.Temp.ExecuteTemplate(w, "Detail", InitStruct.Section)
}

func Categorie(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryCat := r.URL.Query().Get("category")
	lstart := InitStruct.LstCategory(queryCat)
	InitTemp.Temp.ExecuteTemplate(w, "Categorie", lstart)
}

func Search(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryTitle := r.URL.Query().Get("title")
	var lstSearch []InitStruct.Article
	for _, c := range InitStruct.LstArticles {
		if InitStruct.Search(c, queryTitle) {
			lstSearch = append(lstSearch,c)
		}
	}
	InitTemp.Temp.ExecuteTemplate(w, "Search", lstSearch)
}

func Add(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
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
