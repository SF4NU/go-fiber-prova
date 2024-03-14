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

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*") 
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") 
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization") 
		if c.Method() == "OPTIONS" {
			c.Status(fiber.StatusOK)
			return nil
		}
		return c.Next()
	})

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

	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task Task
		if err := c.BodyParser(&task); err != nil {
			return err
		}

		if err := db.Create(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(task)
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		var updatedTask Task
		if err := c.BodyParser(&updatedTask); err != nil {
			return err
		}

		if updatedTask.Description != "" {
			task.Description = updatedTask.Description
		}
		if updatedTask.Completed != task.Completed {
			task.Completed = updatedTask.Completed
		}
		if updatedTask.Deadline != task.Deadline {
			task.Deadline = updatedTask.Deadline
		}

		if err := db.Save(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(task)
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		if err := db.Delete(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusAccepted).SendString("task deleted")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
