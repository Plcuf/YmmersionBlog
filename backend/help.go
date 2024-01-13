package backend

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

// Fonction pour modifié le JSON
func EditJSON(ModifiedArticle []Article) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedArticle)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err := os.WriteFile("JSON/bdd.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}

}

// Fonction pour mettre le JSON dans une struct
func ReadJSON() ([]Article, error) {
	jsonFile, err := os.ReadFile("JSON/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	var jsonData []Article
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

// Fait une liste de dix nombre aléatoire de la taille de la liste
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

// Retoune une liste de blogs de catégorie précise
func LstCategory(category string) []Article {
	var lst []Article
	for _, c := range LstArticles {
		if c.Category == category {
			lst = append(lst, c)
		}
	}
	return lst
}

// Fonction pour rechercher
func Search(word string, s string) bool {
	return strings.Contains(strings.ToLower(word), strings.ToLower(s))
}

// Fonction pour savoir si l'id existe déjà
func IdAlreadyExists(nb int) bool {
	for i := 0; i < len(LstArticles); i++ {
		if LstArticles[i].Id == nb {
			return true
		}
	}
	return false
}

// Fonction pour générer un Id disponible
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

// Fonction pour engregistrer une image et retourne son nom
func InitImg(w http.ResponseWriter, r *http.Request) string {
	//Prend les données ne dépassant cette taille (pout l'image)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	file, handler, errFile := r.FormFile("Image") //Récupère le fichier image
	if errFile != nil {
		http.Error(w, errFile.Error(), http.StatusInternalServerError)
		return ""
	}
	defer file.Close()
	filepath := "./assets/img/" + handler.Filename //Chemin où mettre le fichier
	f, _ := os.Create(filepath)
	defer f.Close()
	io.Copy(f, file) //Met l'image au chemin donnée

	return handler.Filename
}

// Fonction pour récupèrer le mot de passe crypté
func MdpCrypt(Mdp string) string {
	jsonFile, err := os.ReadFile("JSON/login.json") //Récupére les données du JSON
	if err != nil {
		fmt.Println("Error reading", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(jsonFile, &LstUser) //Met dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	hasher := sha256.New()
	hasher.Write([]byte(Mdp))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword // mdp crypter
}
