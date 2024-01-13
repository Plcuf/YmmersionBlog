package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

var err error

// Fonction pour les admins
func Admin(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.Back.User.Admin { //Securisation de la route
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
		return
	}
	InitStruct.Back.Articles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	InitTemp.Temp.ExecuteTemplate(w, "Admin", InitStruct.Back)
}

// Fonction pour ajouter un blog
func Add(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.Back.User.Admin { //Securisation de la route
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "Add", InitStruct.Back)
}

// Fonction treatment de l'ajout de blog
func InitAdd(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.Back.User.Admin { //Securisation de la route
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
		return
	}
	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	//Prend les valeurs demandés
	InitStruct.Section.Category = r.FormValue("category")
	InitStruct.Section.Title = r.FormValue("titre")
	InitStruct.Section.Id = InitStruct.GenerateID() //Je génére un id pas utilisé
	InitStruct.Section.Description = template.HTML(r.FormValue("contains"))
	InitStruct.Section.Author = r.FormValue("Author")
	InitStruct.Section.Introduction = template.HTML(r.FormValue("Introduction"))
	InitStruct.Section.DateCreated = time.Now().Format("2006-01-02")
	InitStruct.Section.Image = InitStruct.InitImg(w, r) //Je sauvegarde l'image dans mon dossier et return le nom du fichier

	InitStruct.LstArticles = append(InitStruct.LstArticles, InitStruct.Section)
	InitStruct.EditJSON(InitStruct.LstArticles) //Met les données dans le JSON
	http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
}

// Fonction pour supprimer un blog
func Suppr(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.Back.User.Admin { //Securisation de la route
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
		return
	}
	InitStruct.LstArticles, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	queryID, errId := strconv.Atoi(r.URL.Query().Get("id")) //Récupére l'Id donné dans le Query string et le met en int
	if errId != nil {
		fmt.Println("Error ID ", errId.Error())
		os.Exit(1)
	}

	queryID-- //Pour que l'index commence à 0
	for _, c := range InitStruct.LstArticles {
		if c.Id == queryID {
			InitStruct.LstArticles = append(InitStruct.LstArticles[:queryID], InitStruct.LstArticles[queryID+1:]...) //Supprime de la liste des blogs
			queryID++
			InitStruct.LstIDSuppr = append(InitStruct.LstIDSuppr, queryID)
			break
		}
	}
	InitStruct.EditJSON(InitStruct.LstArticles) //Met les données dans le JSON
	http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
}

// Fonction pour se déconnecter
func Unlog(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData.Connect = false
	InitStruct.Back.User = InitStruct.Client{"", "", false}
	http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
}

// Fonction pour se connecter
func Login(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData = InitStruct.UserData
	InitTemp.Temp.ExecuteTemplate(w, "Login", InitStruct.Back)
}

// Fonction pour s'inscrire
func Inscription(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData = InitStruct.UserData
	InitTemp.Temp.ExecuteTemplate(w, "inscription", InitStruct.Back)
}

// Fonction treatment pour se connecter
func InitLogin(w http.ResponseWriter, r *http.Request) {
	InitStruct.User.Name = r.FormValue("Nom")
	InitStruct.User.Mdp = InitStruct.MdpCrypt(r.FormValue("Mdp")) //Récupére les données de l'utilisateur

	for _, c := range InitStruct.LstUser {
		if InitStruct.User.Name == c.Name {
			if InitStruct.User.Mdp == c.Mdp {
				InitStruct.UserData.Connect = true //Le connecte
				InitStruct.Back.UserData = InitStruct.UserData
				InitStruct.Back.User = c //Lui met ses droits
				http.Redirect(w, r, InitStruct.Back.UserData.Url, http.StatusMovedPermanently)
				return
			}
		}
	}
	//Sinon reste sur la page login
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}

// Fonction treatment pour se connecter
func InitInscription(w http.ResponseWriter, r *http.Request) {
	InitStruct.User.Name = r.FormValue("Nom")
	InitStruct.User.Mdp = InitStruct.MdpCrypt(r.FormValue("Mdp")) //Récupére les données de l'utilisateur
	InitStruct.User.Admin = false

	for _, c := range InitStruct.LstUser {
		if InitStruct.User.Name == c.Name {
			if InitStruct.User.Mdp == c.Mdp {
				http.Error(w, "Username already exists", http.StatusConflict)
				return
			}
		}
	}

	InitStruct.LstUser = append(InitStruct.LstUser, InitStruct.User) //Ajoute l'utilisateur

	modifiedJSON, errMarshal := json.Marshal(InitStruct.LstUser) //Met la struct en JSON file
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err = os.WriteFile("JSON/login.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}

	InitStruct.UserData.Connect = true
	InitStruct.Back.UserData = InitStruct.UserData
	InitStruct.Back.User = InitStruct.User
	http.Redirect(w, r, InitStruct.Back.UserData.Url, http.StatusMovedPermanently)
}

// Fonction pour la page error 404
func HandleError(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "error", InitStruct.Back)
}
