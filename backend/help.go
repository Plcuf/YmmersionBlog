package backend

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func NbAleatoire(Liste []Article) []int {
	liste := make([]int, 0, 10)
	for len(liste) < 10 {
		randomint := rand.Intn(len(Liste) - 1)
		length := len(liste)
		for i := 0; i < len(liste); i++ {
			if randomint != liste[i] {
				length -= 1
			}
		}
		if length == 0 {
			liste = append(liste, randomint)
		}
	}
	return liste
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

func IdAlreadyExists(nb int) bool {
	for i := 0; i < len(LstArticles); i++ {
		if LstArticles[i].Id == nb {
			return true
		}
	}
	return false
}

func GenerateID() int {
	if !IdAlreadyExists(len(LstArticles) + 1) {
		return len(LstArticles) + 1
	} else {
		t := LstIDSuppr[0]
		if len(LstIDSuppr) > 1 {
			LstIDSuppr = LstIDSuppr[1:]
		} else {
			LstIDSuppr = []int{}
		}
		return t
	}
}