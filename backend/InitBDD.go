package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func EditJSON(ModifiedArticle []Article) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedArticle)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err := os.WriteFile("database/bdd.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}

}

func ReadJSON() ([]Article, error) {
	jsonFile, err := os.ReadFile("database/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	var jsonData []Article
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

func IsInList(lst []string, s string) bool { // on regarde si une lettre est dans la liste ou pas
	for _, c := range lst {
		if string(c) == s {
			return true
		}
	}
	return false
}

func LstCategory(category string) []Article {
	var lst []Article
	for _, c := range LstArticles {
		if c.Category == category {
			lst = append(lst, c)
		}
	}
	return lst
}

func Search(word string, s string) bool {
	return strings.Contains(strings.ToLower(word), strings.ToLower(s))
}

func AddArticle(titre,contains,category,Author,Introduction,Image string,DateCreated int)  {
	var err error
	LstArticles, err = ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	var article Article
	article.Category = category
	article.Title = titre
	article.Id = len(LstArticles) + 1
	article.Description = contains
	article.Author = Author
	article.Introduction = Introduction
	article.DateCreated = DateCreated
	article.Image = Image
	LstArticles = append(LstArticles,article)
	EditJSON(LstArticles)

}