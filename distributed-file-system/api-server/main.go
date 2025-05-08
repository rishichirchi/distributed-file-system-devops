package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

var controlPlaneURL = "http://localhost:5000"
var storageNodes = []string{"http://localhost:6001", "http://localhost:6002", "http://localhost:6003"}

func splitFile(file []byte) [][]byte {
	chunkSize := len(file) / 3
	chunks := [][]byte{}
	for i := 0; i < 3; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == 2 {
			end = len(file)
		}
		chunks = append(chunks, file[start:end])
	}
	return chunks
}

func checksum(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func uploadChunk(chunkID string, data []byte, nodes []string) {
	for _, node := range nodes {
		go func(url string) {
			http.Post(url+"/upload/"+chunkID, "application/octet-stream", bytes.NewReader(data))
		}(node)
	}
}

func main() {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return fiber.ErrBadRequest
		}
		src, _ := file.Open()
		data, _ := io.ReadAll(src)
		chunks := splitFile(data)

		for i, chunk := range chunks {
			chunkID := fmt.Sprintf("%s_chunk_%d", file.Filename, i)
			sum := checksum(chunk)
			uploadChunk(chunkID, chunk, storageNodes)

			meta := map[string]interface{}{
				"chunk_id": chunkID,
				"nodes":    storageNodes,
				"checksum": sum,
			}
			jsonMeta, _ := json.Marshal(meta)
			http.Post(controlPlaneURL+"/register-chunk", "application/json", bytes.NewReader(jsonMeta))
		}
		return c.SendString("Upload complete")
	})

	app.Get("/download/:filename", func(c *fiber.Ctx) error {
		var fileData []byte
		for i := 0; i < 3; i++ {
			chunkID := fmt.Sprintf("%s_chunk_%d", c.Params("filename"), i)
			resp, err := http.Get(controlPlaneURL + "/get-chunk/" + chunkID)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			var meta struct {
				Nodes    []string `json:"nodes"`
				Checksum string   `json:"checksum"`
			}
			json.NewDecoder(resp.Body).Decode(&meta)

			var chunk []byte
			for _, node := range meta.Nodes {
				res, err := http.Get(node + "/chunk/" + chunkID)
				if err != nil || res.StatusCode != 200 {
					continue
				}
				chunk, _ = io.ReadAll(res.Body)
				if checksum(chunk) == meta.Checksum {
					break
				}
			}
			fileData = append(fileData, chunk...)
		}

		return c.Send(fileData)
	})

	app.Listen(":7000")
}