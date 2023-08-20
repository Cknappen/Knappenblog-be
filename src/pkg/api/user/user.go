package user

import (
	"fmt"

	"github.com/Cknappen/Knappenblog-be/src/pkg/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// *HTTP GET* - Gets all the users from the database
func GetAllUsers(c *fiber.Ctx) error {
	db := database.GetLocal[*gorm.DB](c, "db")
	var users []User
	// Retrieve all users
	if result := db.Find(&users); result.Error != nil {
		fmt.Println("Error:", result.Error)

	}
	return c.JSON(users)
}

// *HTTP GET* - Gets a specific user based on user ID
func GetUserById(c *fiber.Ctx) error {
	db := database.GetLocal[*gorm.DB](c, "db")
	var users []User

	// Retrieve all users
	if result := db.Find(&users); result.Error != nil {
		fmt.Println("Error:", result.Error)

	}
	return c.JSON(users)
}

// *HTTP GET* - Gets a specific user based on user name
func GetUserByName(c *fiber.Ctx) error {
	db := database.GetLocal[*gorm.DB](c, "db")
	var users []User

	// Retrieve all users
	if result := db.Find(&users); result.Error != nil {
		fmt.Println("Error:", result.Error)

	}
	return c.JSON(users)
}

// *HTTP POST* - Creates a new user
func CreateUser(c *fiber.Ctx) error {
	db := database.GetLocal[*gorm.DB](c, "db")
	var newUser User
	// Read user from HTTP Post Body.
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body for User",
		})
	}
	// Create the new user record in the database.
	if result := db.Create(&newUser); result.Error != nil {
		fmt.Println("Error:", result.Error)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not write user to database.",
		})
	}
	return c.JSON(newUser)
}

func SetupUserRoutes(db *gorm.DB, group fiber.Router) {
	group.Get("/", GetAllUsers)
	group.Get("/Id", GetUserById)
	group.Get("/Name", GetUserByName)
	group.Post("/", CreateUser)
	// Add similar routes for updating and deleting todos.
}

// Similarly, implement functions for updating and deleting todos.
