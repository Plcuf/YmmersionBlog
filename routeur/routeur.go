package routeur

import (
	ctrl "Ymmersion2/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	//Création des routes pour chaque pages du site
	http.HandleFunc("/index", ctrl.Accueil)
	http.HandleFunc("/jeu/detail", ctrl.Detail)
	http.HandleFunc("/category", ctrl.Category)
	http.HandleFunc("/search", ctrl.Search)
	http.HandleFunc("/admin", ctrl.Admin)
	http.HandleFunc("/add", ctrl.Add)
	http.HandleFunc("/suppr", ctrl.Suppr)
	http.HandleFunc("/add/treatment", ctrl.InitAdd)
	http.HandleFunc("/login", ctrl.Login)
	http.HandleFunc("/login/treatment", ctrl.InitLogin)
	http.HandleFunc("/inscription", ctrl.Inscription)
	http.HandleFunc("/inscription/treatment", ctrl.InitInscription)
	http.HandleFunc("/unlog", ctrl.Unlog)
	http.HandleFunc("/explication", ctrl.Explication)
	http.HandleFunc("/mention", ctrl.Mention)

	//renvoie sur la page d'erreur si la route n'est pas trouvée
	http.HandleFunc("/", ctrl.HandleError)

	//Lien des templates a charger
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	//Le lien d'ou est lancé le serveur
	fmt.Println("(http://localhost:8080/index) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}
