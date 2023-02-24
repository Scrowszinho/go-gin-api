package main

import (
	"github.com/Scrowszinho/go-gin-api/models"
	"github.com/Scrowszinho/go-gin-api/routes"
)

func main() {
	models.Alunos = []models.Student{
		{Name: "Gustavo", RA: 1, CPF: 11122233344},
		{Name: "Ana", RA: 2, CPF: 55566677788},
	}
	routes.HandleRequests()
}
