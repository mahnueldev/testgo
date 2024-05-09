package main

import (
    "fmt"
    "log"
    "example.com/testgo/pkg/routes"
    "example.com/testgo/pkg/database"
    "example.com/testgo/pkg/controllers" // Import the controllers package
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Create db
    database.ConnectDb()

    // Pass the database instance to controllers
    controllers.ConnectDb(database.Database.Db)

    app := fiber.New()

    // Initialize routes
    routes.InitRoutes(app)

    // Start the server
    fmt.Println("Listening on :8080")
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}