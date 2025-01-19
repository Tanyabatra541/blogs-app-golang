package router

import (
	"github.com/Tanyabatra541/blogs-app-golang/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/", controller.BlogUpdate)
	app.Delete("/", controller.BlogDelete)
}