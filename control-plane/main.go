package main

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

type ChunkMetadata struct {
	ChunkID  string   `json:"chunk_id"`
	Nodes    []string `json:"nodes"`
	Checksum string   `json:"checksum"`
}

var chunkMap = make(map[string]ChunkMetadata)
var mu sync.RWMutex

func main() {
	app := fiber.New()

	app.Post("/register-chunk", func(c *fiber.Ctx) error {
		var meta ChunkMetadata
		if err := c.BodyParser(&meta); err != nil {
			return fiber.ErrBadRequest
		}
		mu.Lock()
		chunkMap[meta.ChunkID] = meta
		mu.Unlock()
		return c.SendString("Chunk registered")
	})

	app.Get("/get-chunk/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		mu.RLock()
		meta, ok := chunkMap[id]
		mu.RUnlock()
		if !ok {
			return fiber.ErrNotFound
		}
		return c.JSON(meta)
	})

	app.Listen(":5000")
}