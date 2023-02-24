package models

type Student struct {
	Name string `json:"name"`
	CPF  int    `json:"cpf"`
	RA   int    `json:"ra"`
}

var Alunos []Student
