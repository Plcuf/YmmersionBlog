package main

import (
	r "Ymmersion2/routeur"
	t "Ymmersion2/temps"
)

var err error

func main() {

	t.InitTemplate()
	r.InitServe()

}
