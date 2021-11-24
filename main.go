package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome      string
	Area      string
	Professor string
}

func main() {
	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		panic("erro ao acessar o banco de dados")
	}

	// Código para criar a tabela de Curso no banco de dados
	db.AutoMigrate(&Curso{})

	// Código para criar um novo curso
	db.Create(&Curso{Nome: "Go: Fundamentos de uma aplicação web", Area: "Programacao", Professor: "Guilherme Lima"})

	// Buscar um curso
	var curso Curso
	db.First(&curso, 1) // código para buscar no banco de dados um curso com o id=1

	fmt.Println(curso)
	// Atualizar o curso, mudando o nome
	db.Model(&curso).Update("Nome", "Go: Fundamentos de uma aplicação web com GORM")

	var cursoAtualizado Curso

	db.First(&cursoAtualizado, 1)

	fmt.Println(cursoAtualizado)

	// Código para deletar o curso
	db.Delete(&curso, 1)
}
