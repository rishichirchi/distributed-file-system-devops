package main

import (
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
)

func main() {
	app := fiber.New()
	os.MkdirAll("./chunks", os.ModePerm)

	app.Post("/upload/:chunkID", func(c *fiber.Ctx) error {
		chunkID := c.Params("chunkID")
		data := c.Body()
		err := ioutil.WriteFile("./chunks/"+chunkID, data, 0644)
		if err != nil {
			return err
		}
		return c.SendString("Chunk stored")
	})

	app.Get("/chunk/:chunkID", func(c *fiber.Ctx) error {
		chunkID := c.Params("chunkID")
		data, err := ioutil.ReadFile("./chunks/" + chunkID)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.Send(data)
	})

	app.Listen(":6001")
}