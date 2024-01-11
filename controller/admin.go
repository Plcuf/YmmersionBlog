package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
	}
	InitStruct.Back.Articles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	InitTemp.Temp.ExecuteTemplate(w, "Admin", InitStruct.Back)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
	}

	InitTemp.Temp.ExecuteTemplate(w, "Add", nil)
}

func InitAdd(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
	}
	var err error
	InitStruct.LstArticles, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	InitStruct.Section.Category = r.FormValue("category")
	InitStruct.Section.Title = r.FormValue("titre")
	InitStruct.Section.Id = InitStruct.GenerateID()
	InitStruct.Section.Description = r.FormValue("contains")
	InitStruct.Section.Author = r.FormValue("Author")
	InitStruct.Section.Introduction = r.FormValue("Introduction")
	InitStruct.Section.DateCreated = time.Now().Format("2006-01-02")
	InitStruct.Section.Image = r.FormValue("Image")
	InitStruct.LstArticles = append(InitStruct.LstArticles, InitStruct.Section)
	InitStruct.EditJSON(InitStruct.LstArticles)
	InitTemp.Temp.ExecuteTemplate(w, "Add", nil)
}

func Suppr(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.UserData.Url, http.StatusMovedPermanently)
	}
	queryID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("Error ID ", err.Error())
		os.Exit(1)
	}
	for _, c := range InitStruct.LstArticles {
		if c.Id == queryID {
			InitStruct.LstArticles = append(InitStruct.LstArticles[:queryID], InitStruct.LstArticles[queryID+1:]...)
			InitStruct.LstIDSuppr = append(InitStruct.LstIDSuppr, queryID)
			break
		}
	}
}

func Unlog(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData.Connect = false
	InitStruct.Back.User = InitStruct.Client{"", "",false}
}

func Login(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData = InitStruct.UserData
	InitTemp.Temp.ExecuteTemplate(w, "Login", InitStruct.Back)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	InitStruct.Back.UserData = InitStruct.UserData
	InitTemp.Temp.ExecuteTemplate(w, "inscription", InitStruct.Back)
}

func InitLogin(w http.ResponseWriter, r *http.Request) {
	InitStruct.User.Name = r.FormValue("Nom")
	InitStruct.User.Mdp = r.FormValue("Mdp")
	jsonFile, err := os.ReadFile("database/login.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	err = json.Unmarshal(jsonFile, &InitStruct.LstUser)
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	hasher := sha256.New()
	hasher.Write([]byte(InitStruct.User.Mdp))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	InitStruct.User.Mdp = hashedPassword
	for _, c := range InitStruct.LstUser {
		if InitStruct.User.Name == c.Name {
			if InitStruct.User.Mdp == c.Mdp {
				fmt.Println(InitStruct.UserData)
				InitStruct.User.Admin = c.Admin
				InitStruct.UserData.Connect = true
				InitStruct.Back.UserData = InitStruct.UserData
				InitStruct.Back.User = InitStruct.User
				http.Redirect(w, r, "/index", http.StatusMovedPermanently)
				return
			}
		}
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}

func InitInscription(w http.ResponseWriter, r *http.Request) {
	InitStruct.User.Name = r.FormValue("Nom")
	InitStruct.User.Mdp = r.FormValue("Mdp")
	InitStruct.User.Admin = false

	jsonFile, err := os.ReadFile("database/login.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	err = json.Unmarshal(jsonFile, &InitStruct.LstUser)
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	hasher := sha256.New()
	hasher.Write([]byte(InitStruct.User.Mdp))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	InitStruct.User.Mdp = hashedPassword
	for _, c := range InitStruct.LstUser {
		if InitStruct.User.Name == c.Name {
			if InitStruct.User.Mdp == c.Mdp {
				http.Error(w, "Username already exists", http.StatusConflict)
				return
			}
		}
	}

	InitStruct.LstUser = append(InitStruct.LstUser, InitStruct.User)

	modifiedJSON, errMarshal := json.Marshal(InitStruct.LstUser)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err = os.WriteFile("database/login.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}
	InitStruct.UserData.Connect = true
	InitStruct.Back.UserData = InitStruct.UserData
	InitStruct.Back.User = InitStruct.User
	http.Redirect(w, r, InitStruct.Back.UserData.Url, http.StatusMovedPermanently)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "error", nil)
}
