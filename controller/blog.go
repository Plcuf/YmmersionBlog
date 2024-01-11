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
	//Page d'acceuil avec les recommandations aléatoires
	var lstId []int
	var Recommandation []InitStruct.Article
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	//crée 10 chiffre de 0 au nombre d'articles dispo sur le site et les affiche dans la page d'accueil
	lstId = InitStruct.NbAleatoire(InitStruct.LstArticles)
	for _, i := range lstId {
		Recommandation = append(Recommandation, InitStruct.LstArticles[i])
	}
	InitStruct.UserData.Url = r.URL.String()
	InitStruct.Back.Articles = Recommandation
	//execution du templates acceuil.html
	InitTemp.Temp.ExecuteTemplate(w, "index", InitStruct.Back)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	//prend les datas du json et renvoie un message d'erreur si c'est impossible
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	//prend l'id d'un article dans le json
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
	InitStruct.Back.Articles = []InitStruct.Article{InitStruct.Section}
	//execution du template Detail.html
	InitTemp.Temp.ExecuteTemplate(w, "Detail", InitStruct.Back)
}

func Category(w http.ResponseWriter, r *http.Request) {
	//recupere les categories des data json.
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	queryCat := r.URL.Query().Get("category")
	lstart := InitStruct.LstCategory(queryCat)
	InitStruct.Back.Articles = lstart
	InitStruct.UserData.Url = r.URL.String()
	//execute le temple Category.html
	InitTemp.Temp.ExecuteTemplate(w, "Category", InitStruct.Back)
}

func Search(w http.ResponseWriter, r *http.Request) {
	//Prend le titre des articles enregistré dans le json
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
	//execute le template Search
	InitTemp.Temp.ExecuteTemplate(w, "Search", InitStruct.Back)
}

func Explication(w http.ResponseWriter, r *http.Request) {

	InitTemp.Temp.ExecuteTemplate(w, "explication", InitStruct.Back)
}

func Mention(w http.ResponseWriter, r *http.Request) {

	InitTemp.Temp.ExecuteTemplate(w, "mention", InitStruct.Back)
}
