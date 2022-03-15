package main

import (
	"fmt"

	"github.com/NobunagaYoda/go-rest-api/models"
	"github.com/NobunagaYoda/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}