package main

import (
	"fmt"
	"log"

	"github.com/gmvbr/todo-go/api/config"
	"github.com/gmvbr/todo-go/api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env := config.LoadEnv()

	db, err := database.ConnectDB(env.MongoUrl, env.MongoDb)
	if err != nil {
		log.Fatal(err)
	}
	
    app := fiber.New()
	url := fmt.Sprintf(":%s", env.Port)
    log.Fatal(app.Listen(url))
}