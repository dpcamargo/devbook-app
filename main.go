package main

import (
	"fmt"
	"log"
	"net/http"

	"app/src/router"
	"app/src/router/rotas"
	"app/src/utils"
)

func main() {

	utils.CarregarTemplates()
	
	r := router.Gerar()
	h := rotas.Configurar(r)
	
	fmt.Println("Rodando Webapp")
	log.Fatal(http.ListenAndServe(":5000", h))
}
