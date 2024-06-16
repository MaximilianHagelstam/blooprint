package main

import (
	"fmt"
	"gostarter/internal/database"
	"gostarter/internal/handlers"
	"gostarter/internal/repository"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	logger := log.DefaultLogger()

	db := database.New()

	postHandlerConfig := &handlers.PostHandlerConfig{
		PostRepository: repository.NewPostRepository(db),
		Logger:         logger,
	}

	r := app.Group("/api/v1")
	r.Get("/posts", handlers.GetPostsHandler(postHandlerConfig))
	r.Get("/posts/:id", handlers.GetPostByIDHandler(postHandlerConfig))
	r.Delete("/posts/:id", handlers.DeletePostHandler(postHandlerConfig))

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
