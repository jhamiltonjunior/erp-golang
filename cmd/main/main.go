package main

import (
	"github.com/jhamiltonjunior/cut-url/cmd/main/factor"
	"net/http"
)

func main() {

	controller := factor.MakeURLController()

	http.HandleFunc("/", controller.Create)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
