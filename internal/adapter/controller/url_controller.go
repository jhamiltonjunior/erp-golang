package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"io"
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

type response struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}

type responseError struct {
	status  string
	message string
	err     error
}

func (c *URLController) HandleURL(w http.ResponseWriter, r *http.Request) {

	path := HandleUrl("/url", r.URL.Path)

	methodPath := r.Method + " " + path

	switch methodPath {
	case http.MethodPost + " ":
		c.Create(w, r)
	case http.MethodGet + " /get-by-name":
		c.GetByName(w, r)
	case http.MethodGet + " /get-all":
		c.GetAll(w, r)
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

func (c *URLController) Create(w http.ResponseWriter, r *http.Request) {
	var u url.URL
	var resp response

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		resp = response{
			Status:  "error",
			Message: err.Error(),
		}

		_ = json.NewEncoder(w).Encode(resp)

		return
	}

	err = c.services.Create(u)
	if err != nil {
		resp = response{
			Status:       "error",
			Message:      "Internal Server Error",
			ErrorMessage: err.Error(),
		}

		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = response{
		Status:  "success",
		Message: "created",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)

}

func (c *URLController) GetAll(w http.ResponseWriter, r *http.Request) {
	u, err := c.services.GetAllByUser()
	if err != nil {
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "Internal Server Error",
			"message": "Nao foi possivel Pegar os dados.",
		})
		if err != nil {
			return
		}
	}

	err = json.NewEncoder(w).Encode(&u)
	if err != nil {
		panic(err)
	}
}

func (c *URLController) GetByName(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Description string
	}

	println(r)

	type responseError struct {
		status  string
		message string
		err     error
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r.Body)

	u, err := c.services.GetByName(body.Description)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]error{
			"message": err,
		})

		return
	}

	type response struct {
		status  string
		message string
		u       []url.URL
	}

	_ = json.NewEncoder(w).Encode(u)
}
