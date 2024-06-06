package main

import (
	"log"

	"github.com/Yoginara/ws-yoginara-2024/config"

	"github.com/Yoginara/ws-yoginara-2024/url"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
