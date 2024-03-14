package main

import (

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Deadline    string `json:"deadline"`
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	panic("Error loading .env file")
	// }
	// dbKey := os.Getenv("DB_KEY")

	dsn := "postgres://ryfljouh:dhohPc3uydQ006hRhGsulapMesD4MLFd@dumbo.db.elephantsql.com/ryfljouh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db not connected")
	}

	db.AutoMigrate(&Task{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
