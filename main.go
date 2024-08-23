package main

import (
	"fmt"
	connection "github.com/jhamiltonjunior/cut-url/internal/external/database"
	"github.com/jhamiltonjunior/cut-url/internal/external/factor"
	"net/http"
)

func main() {

	conn := connection.NewMySQLURLRepository("root:0000@tcp(localhost:3306)/cut_url")
	controller := factor.MakeURLController(conn)

	http.HandleFunc("/url", controller.HandleURL)
	http.HandleFunc("/url/", controller.Create)

	fmt.Println("Run ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
