package controller

import (
	InitStruct "Ymmersion2/backend"
	InitTemp "Ymmersion2/temps"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
	}
	InitTemp.Temp.ExecuteTemplate(w, "Admin", nil)
}



func Add(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
	}
	InitTemp.Temp.ExecuteTemplate(w, "add", nil)
}

func AddInit(w http.ResponseWriter, r *http.Request) {
	if !InitStruct.User.Admin {
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
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
		http.Redirect(w, r, InitStruct.User.Url, http.StatusMovedPermanently)
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

func HandleError(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "error", nil)
}