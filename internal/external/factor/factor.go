package factor

import (
	"github.com/jhamiltonjunior/erp-golang/internal/adapter/controller"
	"github.com/jhamiltonjunior/erp-golang/internal/domain/repository"
	connection "github.com/jhamiltonjunior/erp-golang/internal/external/database/mysql"
	"github.com/jhamiltonjunior/erp-golang/internal/external/service"
	"github.com/jhamiltonjunior/erp-golang/internal/usecase"
	"github.com/jhamiltonjunior/erp-golang/internal/usecase/interfaces_usecase"
	"net/http"
	"os"
)

func MakeUserController(conn repository.User, tokenManager interfaces_usecase.Token) *controller.User {
	uc := usecase.NewUserUseCase(conn, tokenManager)
	return controller.NewUserController(*uc)
}

func ServeUser(mux *http.ServeMux) {
	bcrypt := &service.Bcrypt{}
	jwt := &service.JWT{}

	//"root:0000@tcp(localhost:3306)/cut_url"
	conn := connection.NewMySQLUserRepository(os.Getenv("MYSQL_LOCAL_DATABASE"), bcrypt)
	control := MakeUserController(conn, jwt)

	mux.HandleFunc("POST /user", control.Create)
}
