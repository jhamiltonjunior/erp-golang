package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"net/http"
)

type URLController struct {
	services *usecase.URLUseCase
}

func NewURLController(services *usecase.URLUseCase) *URLController {
	return &URLController{
		services: services,
	}
}

func (c *URLController) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := json.NewEncoder(w).Encode(r.Method)

		var data struct {
			OriginalURL string `json:"original_url"`
			UserId      int    `json:"user_id"`
		}

		err = c.services.Create()
		if err != nil {
			return
		}

		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return
		}

		fmt.Println(data)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request method",
		})
		if err != nil {
			panic(err)
		}
	}

}

func (c *URLController) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		url, err := c.services.GetAll()
		if err != nil {
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "Internal Server Error",
				"message": "Nao foi possivel Pegar os dados.",
			})
			if err != nil {
				return
			}
		}

		err = json.NewEncoder(w).Encode(&url)
		if err != nil {
			panic(err)
		}

	default:
		return
	}

}
