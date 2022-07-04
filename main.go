package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/huelet/encode/src/process"
	"github.com/huelet/encode/src/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024 * 1024,
	})

	err := os.MkdirAll("./content/processed", os.ModePerm)
	utils.HandleError(err)
	err = os.MkdirAll("./content/input", os.ModePerm)
	utils.HandleError(err)

	app.Post("/api/process", func(c *fiber.Ctx) error {
		file, err := c.FormFile("video")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.SaveFile(file, fmt.Sprintf("./content/input/%s", file.Filename))
		location := process.Encode(fmt.Sprintf("./content/input/%s", file.Filename), file.Filename)
		resp := process.UploadToAzBlob(location)
		os.Remove(fmt.Sprintf("./content/input/%s", file.Filename))
		os.Remove(location)
		return c.JSON(fiber.Map{"response": resp})
	})

	app.Listen(":3000")
}
