package controllers

import (
	"example.com/testgo/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Define a global variable to hold the database connection
var db *gorm.DB

// Initialize the database connection
func ConnectDb(database *gorm.DB) {
    db = database
}

// GetAllUsers retrieves all users
func GetAllUsers(c *fiber.Ctx) error {
    var users []models.User
    db.Find(&users)
    return c.JSON(users)
}

// GetUser retrieves a single user by ID
func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    db.First(&user, id)
    return c.JSON(user)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
    user := &models.User{}
    if err := c.BodyParser(user); err != nil {
        return err
    }

    db.Create(&user)
    return c.JSON(user)
}

// UpdateUser updates an existing user by ID
func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    db.First(&user, id)

    updatedUser := new(models.User)
    if err := c.BodyParser(updatedUser); err != nil {
        return err
    }

    db.Model(&user).Updates(updatedUser)
    return c.JSON(user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    db.First(&user, id)
    db.Delete(&user)
    return c.SendStatus(fiber.StatusNoContent)
}
