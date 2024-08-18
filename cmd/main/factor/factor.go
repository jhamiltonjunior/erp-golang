package factor

import (
	"github.com/jhamiltonjunior/cut-url/internal/adapter"
	connection "github.com/jhamiltonjunior/cut-url/internal/external"
	"github.com/jhamiltonjunior/cut-url/internal/service"
)

func MakeURLController() *adapter.URLController {
	conn := connection.NewMySQLURLRepository("root:0000@tcp(localhost:3306)/cut_url")
	newService := service.NewURLService(conn)
	return adapter.NewURLController(newService)
}
