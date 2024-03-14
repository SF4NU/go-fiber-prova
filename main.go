package main

import (
	"os"

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
	dsn := "postgres://ryfljouh:dhohPc3uydQ006hRhGsulapMesD4MLFd@dumbo.db.elephantsql.com/ryfljouh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db not connected")
	}
	db.AutoMigrate(&Task{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Todo-api")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Again!")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		var tasks []Task
		db.Find(&tasks)
		return c.Status(fiber.StatusOK).JSON(tasks)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
