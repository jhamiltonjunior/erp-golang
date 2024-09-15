package main

import (
	"fmt"
	"github.com/jhamiltonjunior/erp-golang/internal/external/factor"
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
	factor.ServeUser(mux)

	fmt.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
