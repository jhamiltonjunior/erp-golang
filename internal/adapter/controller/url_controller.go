package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"net/http"
	"strconv"
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
	Data         any    `json:"data"`
}

func (c *URLController) HandleURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodGet:
		c.GetByName(w, r)
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
	var resp response

	id := r.PathValue("user_id")
	idI, err := strconv.Atoi(id)

	fmt.Println(idI)
	fmt.Println(err)

	if err != nil {
		resp = response{
			Status:  "Internal Server Error",
			Message: "No searchable",
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)

		return
	}

	u, err := c.services.GetAllByUser(idI)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "Internal Server Error",
			"message": "Nao foi possivel Pegar os dados.",
		})
	}

	resp = response{
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    u,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (c *URLController) GetByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	description, ok := query["description"]
	if !ok {
		resp := response{
			Status:  "success",
			Message: "No searchable",
		}

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	u, err := c.services.GetByName(description[0])
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]error{
			"message": err,
		})

		return
	}

	resp := response{
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    u,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (c *URLController) Update(w http.ResponseWriter, r *http.Request) {
	var resp response

	var u *url.URL

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		resp = response{
			Status:  "error",
			Message: http.StatusText(http.StatusInternalServerError),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
	}

	err = c.services.Update(u)
	if err != nil {
		resp = response{
			Status:  "error",
			Message: http.StatusText(http.StatusInternalServerError),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = response{
		Status:  "success",
		Message: http.StatusText(http.StatusNoContent),
	}

	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(resp)
}
