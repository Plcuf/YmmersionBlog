package backend

type Site struct {
	Articles []Article
	User     Client
	UserData ClientData
}

type Article struct {
	Title        string `json:"title"`
	Id           int    `json:"id"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Author       string `json:"author"`
	Introduction string `json:"introduction"`
	DateCreated  string `json:"date_created"`
	Image        string `json:"image"`
}

type Client struct {
	Name  string `json:"name"`
	Mdp   string `json:"mdp"`
	Admin bool   `json:"admin"`
}

type ClientData struct {
	Connect bool
	Url     string
}

var Section Article
var LstArticles []Article
var User Client
var UserData ClientData
var LstUser []Client
var Back Site
var LstIDSuppr []int
