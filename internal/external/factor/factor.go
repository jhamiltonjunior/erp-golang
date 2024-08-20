package factor

import (
	"github.com/jhamiltonjunior/cut-url/internal/adapter/controller"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
)

func MakeURLController(conn url_repository.Repository) *controller.URLController {
	newService := usecase.NewURLService(conn)
	return controller.NewURLController(newService)
}
