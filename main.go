package main

import (
	r "Ymmersion2/routeur" //Route vers mes routes
	t "Ymmersion2/temps"   //Route vers mes templates
)

func main() {
	t.InitTemplate() //Initialise mes templates
	r.InitServe()    //Initialise mes routes / assets et lance le serveur
}
