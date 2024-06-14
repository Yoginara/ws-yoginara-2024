package main

import (
	"log"

	"github.com/Yoginara/ws-yoginara-2024/config"

	"github.com/Yoginara/ws-yoginara-2024/url"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/Yoginara/ws-yoginara-2024/docs"
)


// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Yoginara
// @contact.email indra@ulbi.ac.id

// @host ws-yoginara-2024-479b633e2c8b.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
