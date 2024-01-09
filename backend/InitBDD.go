package backend

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var Liste = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

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

func NbAleatoire(Liste []int) []int {
	TailleMax := len(Liste)

	rand.Seed(time.Now().UnixNano())

	NouvelleListe := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		nombreAleatoire := rand.Intn(TailleMax)
		NouvelleListe = append(NouvelleListe, Liste[nombreAleatoire])
	}

	fmt.Println(NouvelleListe)

	return NouvelleListe
}
