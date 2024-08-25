package controller

import (
	"encoding/json"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"net/http"
	"strconv"
)

type URLController struct {
	usecase *usecase.URLUseCase
}

func NewURLController(usecase *usecase.URLUseCase) *URLController {
	return &URLController{
		usecase: usecase,
	}
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
	var resp Response

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		resp = Response{
			Status:  "error",
			Message: err.Error(),
		}

		_ = json.NewEncoder(w).Encode(resp)

		return
	}

	err = c.usecase.Create(u)
	if err != nil {
		resp = Response{
			Status:       "error",
			Message:      "Internal Server Error",
			ErrorMessage: err.Error(),
		}

		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = Response{
		Status:  "success",
		Message: "created",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)

}

func (c *URLController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp Response

	id := r.PathValue("user_id")
	idI, err := strconv.Atoi(id)

	if err != nil {
		resp = Response{
			Status:  "Internal Server Error",
			Message: "No searchable",
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)

		return
	}

	u, err := c.usecase.GetAllByUser(idI)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "Internal Server Error",
			"message": "Nao foi possivel Pegar os dados.",
		})
	}

	resp = Response{
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
		resp := Response{
			Status:  "success",
			Message: "No searchable",
		}

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	u, err := c.usecase.GetByName(description[0])
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]error{
			"message": err,
		})

		return
	}

	resp := Response{
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    u,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (c *URLController) Update(w http.ResponseWriter, r *http.Request) {
	var resp Response

	var u *url.URL

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		resp = Response{
			Status:  "error",
			Message: http.StatusText(http.StatusInternalServerError),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
	}

	err = c.usecase.Update(u)
	if err != nil {
		resp = Response{
			Status:  "error",
			Message: http.StatusText(http.StatusInternalServerError),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = Response{
		Status:  "success",
		Message: http.StatusText(http.StatusNoContent),
	}

	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(resp)
}

func (c *URLController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp Response

	idStr := r.PathValue("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		resp = Response{
			Status:  "error",
			Message: "no searchable",
		}

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	if err = c.usecase.Delete(idInt); err != nil {
		resp = Response{
			Status:  "error",
			Message: http.StatusText(http.StatusInternalServerError),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = Response{
		Status: "success",
	}

	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(resp)
}
