package controller

import (
	"encoding/json"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"net/http"
)

type User struct {
	usecase usecase.User
}

func NewUserController(usecase usecase.User) *User {
	return &User{
		usecase: usecase,
	}
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	var usr entities.User
	var resp Response

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		resp = Response{
			Status:       "error",
			Message:      "Internal Server Error",
			ErrorMessage: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	token, err := user.usecase.Create(usr)
	if err != nil {
		resp = Response{
			Status:       "error",
			Message:      "invalid body",
			ErrorMessage: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp = Response{
		Status:  "success",
		Message: "created",
		Data:    token,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
	return
}
