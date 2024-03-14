package main

import ("github.com/gofiber/fiber/v2"
"os")

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

		port := os.Getenv("PORT")

    app.Listen(":" + port)
}