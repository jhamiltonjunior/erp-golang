package main

import (
	"fmt"
	connection "github.com/jhamiltonjunior/cut-url/internal/external/database"
	"github.com/jhamiltonjunior/cut-url/internal/external/factor"
	"net/http"
)

// export GODEBUG=httpmuxgo121=0
func main() {

	var mux = http.NewServeMux()

	conn := connection.NewMySQLURLRepository("root:0000@tcp(localhost:3306)/cut_url")
	controller := factor.MakeURLController(conn)

	mux.HandleFunc("/url", controller.HandleURL)
	mux.HandleFunc("GET /url/{user_id}", controller.GetAll)
	mux.HandleFunc("PUT /url", controller.Update)
	mux.HandleFunc("DELETE /url/{id}", controller.Delete)

	fmt.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
