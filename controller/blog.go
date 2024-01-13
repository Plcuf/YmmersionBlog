package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Fonction de la page index pour avoir les recommandations aléatoires
func Index(w http.ResponseWriter, r *http.Request) {
	var Recommandation []InitStruct.Article

	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	//Crée 10 chiffre de 0 au nombre d'articles
	lstId := InitStruct.NbAleatoire(InitStruct.LstArticles) //Liste des Id des blogs utilisé dans la recommandation
	for _, i := range lstId {
		Recommandation = append(Recommandation, InitStruct.LstArticles[i])
	}
	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	InitStruct.Back.Articles = Recommandation
	//execution du templates index.html
	InitTemp.Temp.ExecuteTemplate(w, "index", InitStruct.Back)
}

// Fonction de la page du blogs
func Detail(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	//Récupére l'id donné dans le Query string et le convertie en int
	queryID, errId := strconv.Atoi(r.URL.Query().Get("id"))
	if errId != nil {
		fmt.Println("Error ID ", errId.Error())
		os.Exit(1)
	}

	//Prend le blog à l'id donné
	for _, i := range InitStruct.LstArticles {
		if i.Id == queryID {
			InitStruct.Section = i
			break
		}
	}

	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	InitStruct.Back.Articles = []InitStruct.Article{InitStruct.Section}
	//execution du template Detail.html
	InitTemp.Temp.ExecuteTemplate(w, "Detail", InitStruct.Back)
}

// Fonction de la page de la catégorie choisi
func Category(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	queryCat := r.URL.Query().Get("category")  //Récupére la catégorie donné dans le Query string
	lstArticles := InitStruct.LstCategory(queryCat) //Récupére tous les articles de cette catégorie
	InitStruct.Back.Articles = lstArticles
	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	//execute le temple Category.html
	InitTemp.Temp.ExecuteTemplate(w, "Category", InitStruct.Back)
}

// Fonction de la page des résultats d'une recherche
func Search(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	queryTitle := r.URL.Query().Get("title") //Récupére le titre donné dans le Query string
	var lstSearch []InitStruct.Article

	for _, c := range InitStruct.LstArticles {
		if InitStruct.Search(c.Title, queryTitle) {
			lstSearch = append(lstSearch, c) //Met tous les blogs résultant de la recherche
		}
	}

	if len(lstSearch) == 0 { //Oui je triche
		none := InitStruct.Article{Title: "Nothing", Id: 0, Description: "", Category: "", Author: "", Introduction: "", DateCreated: "", Image: ""}
		lstSearch = append(lstSearch, none)
	}
	
	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	InitStruct.Back.Articles = lstSearch
	//execute le template Search
	InitTemp.Temp.ExecuteTemplate(w, "Search", InitStruct.Back)
}

// Fonction pour afficher la page explication
func Explication(w http.ResponseWriter, r *http.Request) {
	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	InitTemp.Temp.ExecuteTemplate(w, "explication", InitStruct.Back)
}

// Fonction pour afficher la page mention
func Mention(w http.ResponseWriter, r *http.Request) {
	InitStruct.UserData.Url = r.URL.String() //Pour savoir où je me situe
	InitTemp.Temp.ExecuteTemplate(w, "mention", InitStruct.Back)
}
