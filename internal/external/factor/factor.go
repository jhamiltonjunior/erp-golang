package factor

import (
	"github.com/jhamiltonjunior/cut-url/internal/adapter/controller"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
	connection "github.com/jhamiltonjunior/cut-url/internal/external/database/mysql"
	"github.com/jhamiltonjunior/cut-url/internal/external/service"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
	"github.com/jhamiltonjunior/cut-url/internal/usecase/interfaces_usecase"
	"net/http"
	"os"
)

func MakeURLController(conn url_repository.Repository) *controller.URLController {
	newService := usecase.NewURLService(conn)
	return controller.NewURLController(newService)
}

func MakeUserController(conn repository.User, tokenManager interfaces_usecase.Token) *controller.User {
	uc := usecase.NewUserUseCase(conn, tokenManager)
	return controller.NewUserController(*uc)
}

func ServeUser(mux *http.ServeMux) {
	bcrypt := &service.Bcrypt{}
	jwt := &service.JWT{}

	conn := connection.NewMySQLUserRepository(os.Getenv("MYSQL_LOCAL_DATABASE"), bcrypt)
	control := MakeUserController(conn, jwt)

	mux.HandleFunc("POST /user", control.Create)
}
