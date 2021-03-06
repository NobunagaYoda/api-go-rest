package main

import (
	"fmt"

	"github.com/NobunagaYoda/go-rest-api/models"
	"github.com/NobunagaYoda/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1 - Joao", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2 - Marcia", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
