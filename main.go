package main

import (
	"fmt"
	"log"
	"net/http"

	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/src/router/rotas"
	"app/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()

	utils.CarregarTemplates()

	r := router.Gerar()
	h := rotas.Configurar(r)

	fmt.Printf("Rodando Webapp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(":5000", h))
}
