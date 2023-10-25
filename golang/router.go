package main

import (
	"db.sampes.puc/controllers"
	"github.com/gofiber/fiber/v2"
)

func initRoutes(app *fiber.App) {
	app.Get("/puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos", controllers.BuscarTodosRobozinhos)
	app.Get("/puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/:id", controllers.BuscarUmRobozinho)
	app.Put("/puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/:id", controllers.AtualizarUmRobozinho)
	app.Post("/puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos", controllers.IncluirUmRobozinho)
	app.Delete("/puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/:id", controllers.RemoverUmRobozinho)
}
