package backend

type Article struct {
	Jeux []struct {
		Name string `json:"name"`
		Article []Section `json:"article"`
	} `json:"jeux"`
}

type Section struct {
	Section struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      string `json:"author"`
		DateCreated int64  `json:"date_created"`
		Image       string `json:"image"`
		Id          int64  `json:"id"`
	} `json:"section"`
}

type Client struct {
	Name  string `json:"name"`
	Mdp   string `json:"mdp"`
	Admin bool   `json:"admin"`
	Url   string
}

var Articles Article
var lstArticles = []Section{}
var User Client
