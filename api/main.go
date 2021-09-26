package main

import (
	"fmt"
	"log"

	"github.com/gmvbr/todo-go/api/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
    app := fiber.New()
	url := fmt.Sprintf(":%s", env.Port)
    log.Fatal(app.Listen(url))
}