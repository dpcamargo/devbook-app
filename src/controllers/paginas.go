package controllers

import (
	"app/src/utils"
	"net/http"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}

func CarregarPaginaCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "cadastro.html", nil)
}
