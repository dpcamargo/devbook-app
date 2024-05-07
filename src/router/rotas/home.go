package rotas

import (
	"app/src/controllers"
	"net/http"
)

var RotaPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
