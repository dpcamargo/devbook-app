package rotas

import (
	"app/src/controllers"
	"net/http"
)

var RotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
