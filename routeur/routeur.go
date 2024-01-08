package routeur

import (
	ctrl "Ymmersion2/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {

	http.HandleFunc("/", ctrl.Accueil)
	http.HandleFunc("/jeu/detail", ctrl.Detail)
	http.HandleFunc("/categorie", ctrl.Categorie)
	http.HandleFunc("/search", ctrl.Search)
	http.HandleFunc("/admin", ctrl.Admin)
	http.HandleFunc("/add", ctrl.Add)
	http.HandleFunc("/error", ctrl.Error)
	
	
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}