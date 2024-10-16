package controllers

import "github.com/gofiber/fiber/v2"

func HelloWold(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}
