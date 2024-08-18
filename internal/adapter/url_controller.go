package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/service"
	"net/http"
)

type URLController struct {
	services *service.URLService
}

func NewURLController(services *service.URLService) *URLController {
	return &URLController{
		services: services,
	}
}

func (c *URLController) Create(w http.ResponseWriter, r *http.Request) {

	var data struct {
		OriginalURL string `json:"original_url"`
		UserId      int    `json:"user_id"`
	}

	err := c.services.Create()
	if err != nil {
		return
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}

	fmt.Println(data)
}
