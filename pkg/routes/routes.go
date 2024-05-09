package routes

import (
	"example.com/testgo/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
    api := app.Group("/api")

    // User routes
    userRoutes := api.Group("/users")
    userRoutes.Get("/", controllers.GetAllUsers)
    userRoutes.Get("/:id", controllers.GetUser)
    userRoutes.Post("/", controllers.CreateUser)
    userRoutes.Put("/:id", controllers.UpdateUser)
    userRoutes.Delete("/:id", controllers.DeleteUser)
}
