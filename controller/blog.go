package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var err error

func Accueil(w http.ResponseWriter, r *http.Request) {
	var lstId []int
	var Recommandation []InitStruct.Article
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	lstId = InitStruct.NbAleatoire(InitStruct.LstArticles)
	for _, i := range lstId {
		Recommandation = append(Recommandation, InitStruct.LstArticles[i])
	}
	InitStruct.UserData.Url = r.URL.String()
	InitStruct.Back.Articles = Recommandation
	InitTemp.Temp.ExecuteTemplate(w, "index", InitStruct.Back)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryID, errId := strconv.Atoi(r.URL.Query().Get("id"))
	if errId != nil {
		fmt.Println("Error ID ", errId.Error())
		os.Exit(1)
	}
	for _, i := range InitStruct.LstArticles {
		if i.Id == queryID {
			InitStruct.Section = i
			break
		}
	}
	InitStruct.UserData.Url = r.URL.String()
	InitStruct.Section.Description = InitStruct.Replace(InitStruct.Section.Description)
	InitStruct.Back.Articles = []InitStruct.Article{InitStruct.Section}
	fmt.Println(InitStruct.Back.Articles[0].Category)
	InitTemp.Temp.ExecuteTemplate(w, "Detail", InitStruct.Back)
}

func Category(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryCat := r.URL.Query().Get("category")
	lstart := InitStruct.LstCategory(queryCat)
	InitStruct.Back.Articles = lstart
	InitStruct.UserData.Url = r.URL.String()
	InitTemp.Temp.ExecuteTemplate(w, "Category", InitStruct.Back)
}

func Search(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryTitle := r.URL.Query().Get("title")
	var lstSearch []InitStruct.Article
	for _, c := range InitStruct.LstArticles {
		if InitStruct.Search(c.Title, queryTitle) {
			lstSearch = append(lstSearch, c)
		}
	}
	InitStruct.UserData.Url = r.URL.String()
	InitStruct.Back.Articles = lstSearch
	InitTemp.Temp.ExecuteTemplate(w, "Search", InitStruct.Back)
}
