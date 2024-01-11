package routeur

import (
	ctrl "Ymmersion2/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {

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


	http.HandleFunc("/", ctrl.HandleError)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/index) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}
