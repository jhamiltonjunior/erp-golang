package main

import (
	"fmt"
	connection "github.com/jhamiltonjunior/cut-url/internal/external/database/mysql"
	"github.com/jhamiltonjunior/cut-url/internal/external/factor"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// export GODEBUG=httpmuxgo121=0
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var mux = http.NewServeMux()

	conn := connection.NewMySQLURLRepository("root:0000@tcp(localhost:3306)/cut_url")
	controller := factor.MakeURLController(conn)

	mux.HandleFunc("/url", controller.HandleURL)
	mux.HandleFunc("GET /url/{user_id}", controller.GetAll)
	mux.HandleFunc("PUT /url", controller.Update)
	mux.HandleFunc("DELETE /url/{id}", controller.Delete)

	factor.ServeUser(mux)

	fmt.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
