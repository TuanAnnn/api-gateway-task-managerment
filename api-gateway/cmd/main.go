package main

import (
	"fmt"
	"log"

	"github.com/TuanAnnn/api-gateway/pkg/common/config"
	user "github.com/TuanAnnn/api-gateway/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	user.RegisterRoutes(app, c)

	err = app.Listen(fmt.Sprintf(":%s", c.Port))

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}
}
