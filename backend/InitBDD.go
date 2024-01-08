package backend

import (
	"encoding/json"
	"fmt"
	"os"
)

func EditJSON(ModifiedArticle Article) {

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

func ReadJSON() (Article, error) {
	jsonFile, err := os.ReadFile("database/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	var jsonData Article
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}
